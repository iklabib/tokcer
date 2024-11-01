package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/iklabib/tokcer/memo"
	"github.com/iklabib/tokcer/tiktok"
	"github.com/labstack/echo/v4"
)

var tk *tiktok.Tiktok = tiktok.NewTiktok()

func main() {
	e := echo.New()

	e.POST("/search", func(c echo.Context) error {
		var req VideoSearchRequest
		if err := c.Bind(&req); err != nil {
			log.Printf("search: failed to read request body %s", err.Error())
			return c.JSON(400, "bad request")
		}

		key := strings.TrimSpace(req.Keywords)
		vd := memo.LoadSearch(key)
		if vd == nil {
			vd = tk.VideoSearch(key)
			if vd != nil {
				memo.AddSearch(key, vd)
			} else {
				log.Printf("video search failed keywords: '%s' type: '%s'", key, req.SearchType)
				return c.JSON(500, "interal server error")
			}
		}

		results := vd.LoadAll()

		return c.JSON(200, results)
	})

	e.POST("/video", func(c echo.Context) error {
		var req VideoManifestRequest
		if err := c.Bind(&req); err != nil {
			log.Printf("video: failed to read request body %s", err.Error())
			return c.JSON(400, "bad request")
		}

		vinfo, err := tk.GetVideoInfo(req.Url)
		if err != nil {
			log.Printf("video: %s", err.Error())
			return c.JSON(500, "failed to fetch video information")
		}

		return c.JSON(200, vinfo)
	})

	e.GET("/stream", func(c echo.Context) error {
		// var req VideoStreamRequest
		// if err := c.Bind(&req); err != nil {
		// 	return c.JSON(http.StatusBadRequest, "bad request")
		// }

		user := c.QueryParam("u")
		videoId := c.QueryParam("id")
		url := fmt.Sprintf("https://www.tiktok.com/%s/video/%s", user, videoId)

		vs, err := tk.StreamVideo(url)
		if err != nil {
			log.Printf("stream: %s", err.Error())
			return c.JSON(500, "failed to stream video")
		}

		return c.Stream(200, "video/"+vs.Ext, vs.Video)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
