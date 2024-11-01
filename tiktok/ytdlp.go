package tiktok

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func streamVideo(videoUrl string) (*VideoStream, error) {
	vi, err := getVideoInfo(videoUrl)
	if err != nil {
		return nil, err
	}

	s, err := getVideoStream(vi)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// hacky, but reliable
func getVideoInfo(url string) (*VideoInfo, error) {
	cmd := exec.Command("yt-dlp", "-j", url)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running yt-dlp: %w", err)
	}

	var videoInfo VideoInfo
	if err := json.Unmarshal(output, &videoInfo); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &videoInfo, nil
}

func getVideoStream(videoInfo *VideoInfo) (*VideoStream, error) {
	format := pickVideo(videoInfo.Formats)
	headers := format.HTTPHeaders

	cookies := parseCookies(format.Cookies)

	req, err := http.NewRequest("GET", format.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	for key, value := range cookies {
		req.AddCookie(&http.Cookie{
			Name:  key,
			Value: value,
		})
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	return &VideoStream{Ext: format.Ext, Video: resp.Body}, nil
}

func parseCookies(cookiesStr string) map[string]string {
	cookies := make(map[string]string)
	for _, cookie := range strings.Split(cookiesStr, "; ") {
		if !strings.Contains(cookie, "=") {
			continue
		}

		kv := strings.SplitN(cookie, "=", 2)
		if strings.TrimSpace(kv[0]) == "" || strings.TrimSpace(kv[1]) == "" {
			continue
		}
		cookies[kv[0]] = kv[1]
	}
	return cookies
}

// pick highest quality unwatermarked video format
func pickVideo(formats []Format) Format {
	item := formats[0]
	backseat := formats[0]

	picked := false
	for _, current := range formats {
		if current.FormatNote != "watermarked" {
			continue
		}

		// favor h264 for availability, quality may suffer
		if current.Quality > item.Quality {
			if current.Vcodec == "h264" {
				picked = true
				item = current
			} else {
				backseat = current
			}
		}
	}

	// no h264?
	if !picked {
		return backseat
	}

	return item
}

// "https://www.tiktok.com/@gayscosmo/video/7415341992571374854"
