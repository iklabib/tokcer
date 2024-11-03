package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/iklabib/tokcer/memo"
	"github.com/iklabib/tokcer/tiktok"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// To anyone reading this code:
// I apologize that you have to witness this atrocity
// a barely functional solution, duct-taped together on a weekend

var tk *tiktok.Tiktok = tiktok.NewTiktok()

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	// periodically clean cache
	memo.StartCacheCleaner(10 * time.Minute)

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

	e.GET("/video/:user/:id", func(c echo.Context) error {
		user := c.Param("user")
		videoId := c.Param("id")
		url := fmt.Sprintf("https://www.tiktok.com/%s/video/%s", user, videoId)

		vinfo := memo.LoadVideoInfo(url)
		if vinfo == nil {
			vi, err := tk.GetVideoInfo(url)
			if err != nil {
				log.Printf("video: %s", err.Error())
				return c.JSON(500, "failed to fetch video information")
			}
			memo.AddVideo(url, vi)

			vinfo = vi
		}

		return c.JSON(200, vinfo)
	})

	e.HEAD("/stream/:user/:id", func(c echo.Context) error {
		user := c.Param("user")
		videoId := c.Param("id")
		url := fmt.Sprintf("https://www.tiktok.com/%s/video/%s", user, videoId)

		vs := memo.LoadStream(url)
		if vs == nil {
			st, err := tk.StreamVideo(url)
			if err != nil {
				log.Printf("stream: %s", err.Error())
				return c.JSON(500, "failed to stream video")
			}

			vs = st
		}

		c.Response().Header().Set("Content-Type", "video/"+vs.Ext)

		return c.NoContent(200)
	})

	e.GET("/stream/:user/:id", func(c echo.Context) error {
		// TODO: range support
		user := c.Param("user")
		videoId := c.Param("id")
		url := fmt.Sprintf("https://www.tiktok.com/%s/video/%s", user, videoId)

		// well, video player should hit HEAD fist
		// but just check them just in case
		vs := memo.LoadStream(url)
		if vs == nil {
			st, err := tk.StreamVideo(url)
			if err != nil {
				log.Printf("stream: %s", err.Error())
				return c.JSON(500, "failed to stream video")
			}

			vs = st
		}

		defer memo.DeleteStream(url)

		c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", vs.ContentLength))

		return c.Stream(200, "video/"+vs.Ext, vs.Video)
	})

	host := os.Getenv("TOKCER_HOST")
	if host != "" {
		e.Logger.Fatal(e.Start(host))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}
