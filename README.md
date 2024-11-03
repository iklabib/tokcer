# Tokcer
Using go-rod & yt-dlp to extract information from TikTok.

Requirment
- Recent version of Go
- Bun or Node js
- Chrome or Chromium
- [yt-dlp](https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#installation) in PATH

```shell
$ git clone https://github.com/iklabib/tokcer
$ cd tokcer
$ VITE_API_URL="http://localhost:1323"
$ TOKCER_HOST=$VITE_API_URL
$ go run tokcer/cmd/*.go &
$ cd web
$ bun i
$ bun run dev
```

> [!IMPORTANT]  
> The current stage of Tokcer requires using Go Rod in user mode as a workaround for TikTok's anti-bot. You will need to close all Chrome or Chromium instances on your machine for it to work.
