package main

type VideoSearchRequest struct {
	Keywords   string `json:"keywords"`
	SearchType string `json:"search_type"`
}

type VideoManifestRequest struct {
	User    string `json:"user"`
	VideoId string `json:"video_id"`
}
