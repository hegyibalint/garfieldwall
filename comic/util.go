package comic

import "image"

func TotalHeight(strips []image.Image) int {
	height := 0
	for _, strip := range strips {
		height += strip.Bounds().Dy()
	}
	return height
}

func TotalWidth(strips []image.Image) int {
	width := 0
	for _, strip := range strips {
		width += strip.Bounds().Dx()
	}
	return width
}
