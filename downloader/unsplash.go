package downloader

import (
	"fmt"
	"strconv"
	"strings"
	"unsplash-background/config"
)

func UnsplashDownload() string {
	url := "https://source.unsplash.com/"
	if config.Unsplash.SelectRandom {
		url = url + "random/" + strconv.Itoa(config.Unsplash.Resolution.Width) + "x" + strconv.Itoa(config.Unsplash.Resolution.Height)
	} else {
		url = url + strconv.Itoa(config.Unsplash.Resolution.Width) + "x" + strconv.Itoa(config.Unsplash.Resolution.Height) + "/?" + strings.Join(config.Unsplash.Search, ",")
	}
	fmt.Println("Start downloading from ", url)

	err := DownloadFile(url)
	if err != nil {
		panic(fmt.Sprintf("Error cannot download from: %s", url))
	}

	return tmpPath
}
