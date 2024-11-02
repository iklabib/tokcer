package tiktok

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/stealth"
)

func (tk *Tiktok) VideoSearch(keywords string) *VideoSearch {
	return NewVideoSearch(tk.browser, keywords)
}

func NewVideoSearch(browser *rod.Browser, keywords string) *VideoSearch {
	searchURL := fmt.Sprintf("https://www.tiktok.com/search/video?q=%s&t=%d", url.QueryEscape(keywords), time.Now().Unix())

	page := stealth.MustPage(browser)

	randomDelay()

	page.MustNavigate(searchURL)

	timeout := 1500 * time.Millisecond
	page.Timeout(timeout).Element(`div[data-e2e="search_video-item"]`)

	return &VideoSearch{
		browser:  browser,
		page:     page,
		keywords: keywords,
		index:    map[string]bool{},
	}
}

func (vs *VideoSearch) Close() {
	vs.page.MustClose()
}

// return differences between call
func (vs *VideoSearch) Load() []VideoSearchItem {
	vs.page.Keyboard.Press(input.End)
	r := rand.Int63n(1500-500) + 500
	wait := time.Duration(r) * time.Millisecond
	vs.page.WaitIdle(wait)

	items := vs.parseVideoSearchPage()
	var newItems []VideoSearchItem

	for _, item := range items {
		if _, ok := vs.index[item.Url]; ok {
			continue
		}
		vs.index[item.Url] = true
		newItems = append(newItems, item)
	}

	return newItems
}

// just like Load, but return results as is
func (vs *VideoSearch) LoadAll() []VideoSearchItem {
	vs.page.Keyboard.Press(input.End)

	randomDelay()

	return vs.parseVideoSearchPage()
}

func (vs *VideoSearch) parseVideoSearchPage() []VideoSearchItem {
	var items VideoSearchItems
	script := `(() => {
		try {
			const videos = []
			for (const item of document.querySelectorAll('div[data-e2e="search_video-item"]')) {
				const p = item.parentElement
				const videoItem = p.querySelector('div[data-e2e="search_video-item"]')
				const videoDesc = Array.from(p.querySelector('div[data-e2e="search-card-video-caption"] h1').childNodes)
					.map(el => el.textContent)
					.join('')
				const video = {
					url: videoItem.querySelector('a').href,
					desc: videoDesc,
					coverAlt: videoItem.querySelector('img').alt,
					cover: videoItem.querySelector('img').src,
				}
				videos.push(video)
			}
			return {videos: videos}
		} catch (e) {
			return {videos: []}
		}
	})`

	o, err := vs.page.Eval(script)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	err = o.Value.Unmarshal(&items)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return items.Videos
}
