package main

type VideoSearchRequest struct {
	Keywords   string `json:"keywords"`
	SearchType string `json:"search_type"`
}

type VideoManifestRequest struct {
	Url string `json:"url"`
}
