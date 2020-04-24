package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var tmpPath string = "/tmp/downloaded_wallpaper"

// WriteCounter counts the number of bytes written to it. By implementing the write method,
// it is of the io.Writer interface and we can pass this into io.TeeReader()
// Every write to this writer, will print the progress of the file write.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	return n, nil
}

func DownloadFile(url string) error {
	// Create the file with .tmp extension, so that we won't overwrite a
	// file until it's downloaded fully
	tmpPath = tmpPath + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg"
	out, err := os.Create(tmpPath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create our bytes counter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Println()

	// Rename the tmp file back to the original file
	err = os.Rename(tmpPath+".tmp", tmpPath)
	if err != nil {
		return err
	}

	return nil
}
