// tiktok.go
package tiktok

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

type Tiktok struct {
	browser *rod.Browser
}

func NewTiktok() *Tiktok {
	// the only workaround against TikTok's anti bot is using user mode
	wsURL := launcher.NewUserMode().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(wsURL).MustConnect().NoDefaultDevice()
	return &Tiktok{browser: browser}
}

func (tk *Tiktok) Close() {
	tk.browser.MustClose()
}

func randomDelay() {
	delay := time.Duration(500+rand.Intn(1500)) * time.Millisecond
	time.Sleep(delay)
}

func (tk *Tiktok) GetVideoInfo(url string) (*VideoInfoMin, error) {
	page := stealth.MustPage(tk.browser)
	defer page.Close()

	page.MustNavigate(url)

	err := page.WaitLoad()
	if err != nil {
		return nil, fmt.Errorf("timeout waiting for page load: %v", err)
	}

	_, err = page.Element("#__UNIVERSAL_DATA_FOR_REHYDRATION__")
	if err != nil {
		return nil, fmt.Errorf("failed to find data element: %v", err)
	}

	randomDelay()

	var videoInfo VideoInfoMin
	script := `() => {
		try {
			const vd = JSON.parse(document.getElementById("__UNIVERSAL_DATA_FOR_REHYDRATION__").text).__DEFAULT_SCOPE__["webapp.video-detail"]["itemInfo"]["itemStruct"]
			const relatedVideoQuery = document.querySelectorAll("div[class*='DivItemContainer']")
			const relatedVideos = []
			for (const item of relatedVideoQuery) {
				const current = item.querySelector('a[title]')
				if (current) {
					const username = item.querySelector('div[class*="DivAuthor"]').textContent
					const img = current.querySelector('img')?.src ?? ""
					relatedVideos.push({
						thumbnail: img,
						url: current.href,
						username: username,
						title: current.title,
					})
				}
			}
			return {
				desc: vd.desc,
				author: vd.author,
				statsV2: vd.statsV2,
				relatedVideos: relatedVideos,
			}
		} catch (e) {
			return {};
		}
	}`
	err = page.MustEval(script).Unmarshal(&videoInfo)

	if err != nil {
		return nil, fmt.Errorf("failed to extract video info: %v", err)
	}

	return &videoInfo, nil
}

func (tk *Tiktok) StreamVideo(url string) (*VideoStream, error) {
	return streamVideo(url)
}
