package tiktok

import (
	"io"

	"github.com/go-rod/rod"
)

type VideoSearch struct {
	browser  *rod.Browser
	page     *rod.Page
	keywords string
	index    map[string]bool
}

type VideoSearchItems struct {
	Videos []VideoSearchItem `json:"videos"`
}

type VideoSearchItem struct {
	Url        string `json:"url"`
	Desc       string `json:"desc"`
	CoverAlt   string `json:"coverAlt"`
	Cover      string `json:"cover"`
	Username   string `json:"username"`
	UserAvatar string `json:"userAvatar"`
}

type RelatedVideo struct {
	Url       string `json:"url"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Username  string `json:"username"`
}

type VideoInfoMin struct {
	Desc          string         `json:"desc"`
	RelatedVideos []RelatedVideo `json:"relatedVideos"`
	Author        struct {
		ID           string `json:"id"`
		ShortID      string `json:"shortId"`
		UniqueID     string `json:"uniqueId"`
		Nickname     string `json:"nickname"`
		AvatarLarger string `json:"avatarLarger"`
		AvatarMedium string `json:"avatarMedium"`
		AvatarThumb  string `json:"avatarThumb"`
		Signature    string `json:"signature"`
	} `json:"author"`
	StatsV2 struct {
		DiggCount    string `json:"diggCount"`
		ShareCount   string `json:"shareCount"`
		CommentCount string `json:"commentCount"`
		PlayCount    string `json:"playCount"`
		CollectCount string `json:"collectCount"`
		RepostCount  string `json:"repostCount"`
	} `json:"statsV2"`
}

type Music struct {
	Title       string `json:"title"`
	CoverLarge  string `json:"coverLarge"`
	CoverMedium string `json:"coverMedium"`
	CoverThumb  string `json:"coverThumb"`
	AuthorName  string `json:"authorName"`
	Duration    int    `json:"duration"`
}

type Video struct {
	ID            string `json:"id"`
	Height        int    `json:"height"`
	Width         int    `json:"width"`
	Duration      int    `json:"duration"`
	Ratio         string `json:"ratio"`
	Cover         string `json:"cover"`
	Bitrate       int    `json:"bitrate"`
	EncodedType   string `json:"encodedType"`
	Format        string `json:"format"`
	VideoQuality  string `json:"videoQuality"`
	EncodeUserTag string `json:"encodeUserTag"`
	CodecType     string `json:"codecType"`
	Definition    string `json:"definition"`
}

type VideoStream struct {
	Ext           string        `json:"ext"`
	Video         io.ReadCloser `json:"video"`
	ContentLength int64         `json:"content_length"`
}

type VideoInfo struct {
	ID      string   `json:"id"`
	Formats []Format `json:"formats"`
}

type Format struct {
	URL         string            `json:"url"`
	HTTPHeaders map[string]string `json:"http_headers"`
	Cookies     string            `json:"cookies"`
	Ext         string            `json:"ext"`
	Vcodec      string            `json:"vcodec"`
	Acodec      string            `json:"acodec"`
	FormatNote  string            `json:"format_note"`
	Protocol    string            `json:"protocol"`
	Resolution  *interface{}      `json:"resolution,omitempty"`
	AspectRatio *interface{}      `json:"aspect_ratio,omitempty"`
	VideoExt    string            `json:"video_ext"`
	AudioExt    string            `json:"audio_ext"`
	Quality     int               `json:"quality,omitempty"`
	Width       int               `json:"width,omitempty"`
	Height      int               `json:"height,omitempty"`
}
