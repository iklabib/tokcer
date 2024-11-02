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

	randomDelay()

	page.MustNavigate(url)

	err := page.WaitLoad()
	if err != nil {
		return nil, fmt.Errorf("timeout waiting for page load: %v", err)
	}

	randomDelay()
	_, err = page.Element("#__UNIVERSAL_DATA_FOR_REHYDRATION__")
	if err != nil {
		return nil, fmt.Errorf("failed to find data element: %v", err)
	}

	// Execute JavaScript with error handling
	var videoInfo VideoInfoMin
	script := `() => {
		try {
			const vd = JSON.parse(document.getElementById("__UNIVERSAL_DATA_FOR_REHYDRATION__").text).__DEFAULT_SCOPE__["webapp.video-detail"]["itemInfo"]["itemStruct"]
			return {
				desc: vd.desc,
				author: vd.author,
				// music: vd.music,
				// video: vd.video,
				statsV2: vd.statsV2,
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
