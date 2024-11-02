package memo

import (
	"log"
	"sync"
	"time"

	"github.com/iklabib/tokcer/tiktok"
)

// roll your own cache is not the brightest idea
// but here we are

var (
	searchCache      = map[string]*tiktok.VideoSearch{}
	videoCache       = map[string]*tiktok.VideoInfoMin{}
	videoStream      = map[string]*tiktok.VideoStream{}
	searchCacheMutex = &sync.Mutex{}
	videoCacheMutex  = &sync.Mutex{}
	videoStreamMutex = &sync.Mutex{}
)

func LoadSearch(key string) *tiktok.VideoSearch {
	searchCacheMutex.Lock()
	defer searchCacheMutex.Unlock()
	return searchCache[key]
}

func AddSearch(key string, vd *tiktok.VideoSearch) {
	searchCacheMutex.Lock()
	defer searchCacheMutex.Unlock()
	searchCache[key] = vd
}

func LoadVideoInfo(key string) *tiktok.VideoInfoMin {
	videoCacheMutex.Lock()
	defer videoCacheMutex.Unlock()
	return videoCache[key]
}

func AddVideo(key string, vi *tiktok.VideoInfoMin) {
	videoCacheMutex.Lock()
	defer videoCacheMutex.Unlock()
	videoCache[key] = vi
}

// keep in mine that you may only read video body ONCE
// delete the element after read
func LoadStream(key string) *tiktok.VideoStream {
	videoStreamMutex.Lock()
	defer videoStreamMutex.Unlock()
	return videoStream[key]
}

func AddStream(key string, st *tiktok.VideoStream) {
	videoStreamMutex.Lock()
	defer videoStreamMutex.Unlock()
	videoStream[key] = st
}

func DeleteStream(key string) {
	videoStreamMutex.Lock()
	defer videoStreamMutex.Unlock()
	delete(videoStream, key)
}

func StartCacheCleaner(d time.Duration) {
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for range ticker.C {
			cleanCaches()
		}
	}()
}

func cleanCaches() {
	searchCacheMutex.Lock()
	searchCache = map[string]*tiktok.VideoSearch{}
	searchCacheMutex.Unlock()

	videoCacheMutex.Lock()
	videoCache = map[string]*tiktok.VideoInfoMin{}
	videoCacheMutex.Unlock()

	videoStreamMutex.Lock()
	videoStream = map[string]*tiktok.VideoStream{}
	videoStreamMutex.Unlock()

	log.Println("Caches cleaned")
}
