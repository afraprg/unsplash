package main

import (
	"github.com/martinlindhe/notify"
	"github.com/reujab/wallpaper"
	"time"
	"unsplash-background/config"
	"unsplash-background/downloader"
)

func main() {

	ticker := time.NewTicker(time.Duration(config.Unsplash.Interval) * time.Hour)

	err := downloadAndSetBackground()
	if err != nil {
		return
	}

	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := downloadAndSetBackground()
			if err != nil {
				return
			}
		}
	}

	done <- true
}

func downloadAndSetBackground() error {
	wallpaperPath := downloader.UnsplashDownload()
	notify.Notify("Unsplash", "Unsplash", "Background Changed", "")
	err := wallpaper.SetFromFile(wallpaperPath)

	return err
}
