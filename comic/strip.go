package comic

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"log"
	"net/http"
	"time"
)

var (
	baseURL = "https://d1ejxu6vysztl5.cloudfront.net/comics/garfield/%d/%d-%02d-%02d.gif"
)

func DownloadStrip(width int, time time.Time) (strip image.Image, err error) {
	url := fmt.Sprintf(baseURL, time.Year(), time.Year(), time.Month(), time.Day())

	log.Print(fmt.Sprintf("Downloading image %s", time))
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	log.Print(fmt.Sprintf("Decoding image %s", time))
	strip, err = gif.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	// Calculate the target height size from the ratio:
	// (actual width / target width)
	scaleRatio := (float64(width) / float64(strip.Bounds().Dx()))
	targetHeight := float64(strip.Bounds().Dy()) * scaleRatio
	resizedStrip := resize.Resize(uint(width), uint(targetHeight), strip, resize.Bicubic)

	return resizedStrip, nil
}
