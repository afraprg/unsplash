package main

import (
	"github.com/martinlindhe/notify"
	"time"
	"unsplash-background/config"
	"unsplash-background/downloader"
	"unsplash-background/setter"
)

func main() {

	ticker := time.NewTicker(time.Duration(config.Unsplash.Interval) * time.Hour)

	err := downloadAndSetBackground()
	if err != nil {
		return
	}

	for {
		select {
		case <-ticker.C:
			err := downloadAndSetBackground()
			if err != nil {
				return
			}
		}
	}
}

func downloadAndSetBackground() error {
	wallpaperPath := downloader.UnsplashDownload()
	notify.Notify("Unsplash", "Unsplash", "Background Changed", "")
	err := setter.SetFromFile(wallpaperPath)

	return err
}
