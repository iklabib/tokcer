package memo

import "github.com/iklabib/tokcer/tiktok"

var searchCache map[string]*tiktok.VideoSearch = map[string]*tiktok.VideoSearch{}

func LoadSearch(key string) *tiktok.VideoSearch {
	v, ok := searchCache[key]
	if !ok {
		return nil
	}

	return v
}

func AddSearch(key string, vd *tiktok.VideoSearch) {
	searchCache[key] = vd
}
