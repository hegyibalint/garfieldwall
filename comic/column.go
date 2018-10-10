package comic

import (
	"image"
	"image/color"
	"image/draw"
	"log"
)

func CreateColumn(index, width, height int, timeRange TimeRange) (column image.Image, err error) {
	var strips []image.Image

	log.Printf("Creating column #%d\n", index+1)

	for TotalHeight(strips) < height {
		stripDate, err := GenerateRandomDate(timeRange)
		if err != nil {
			return nil, err
		}

		strip, err := DownloadStrip(width, stripDate)
		if err != nil {
			return nil, err
		}
		strips = append(strips, strip)
	}

	// TODO: replace this with the complete array when optimizeSize is implemented
	selectedComics := optimizeSize(strips[:len(strips)-1], height)
	column = mergeStripsIntoColumn(width, height, selectedComics)
	return column, nil
}

func optimizeSize(strips []image.Image, height int) []image.Image {
	// TODO: implement this
	return strips
}

func mergeStripsIntoColumn(width, height int, strips []image.Image) image.Image {
	column := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(
		column,
		column.Bounds(),
		&image.Uniform{color.White},
		image.ZP,
		draw.Src)

	columnY := (height - TotalHeight(strips)) / 2
	for _, strip := range strips {
		targetRect := image.Rect(
			0,
			columnY,
			strip.Bounds().Dx(),
			columnY + strip.Bounds().Dy())
		draw.Draw(
			column,
			targetRect,
			strip,
			image.ZP,
			draw.Src)
		columnY += strip.Bounds().Dy()
	}

	return column
}
