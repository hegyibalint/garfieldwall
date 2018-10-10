package comic

import (
	"bufio"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

var (
	comicWidth = 1200
)

func CreateWall(width, height int, scale float64, timeRange TimeRange) (wall image.Image, err error) {
	targetStripWidth := int(float64(comicWidth) / scale)
	requiredColumns := width / targetStripWidth

	log.Printf("Comic strip will be %d px wide\n", targetStripWidth)
	log.Printf("Image will contains %d columns\n", requiredColumns)

	var columns []image.Image
	for i := 0; i < requiredColumns; i++ {
		column, err := CreateColumn(i, targetStripWidth, height, timeRange)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	wall = mergeColumnsIntoWall(width, height, columns)
	return wall, nil
}

func mergeColumnsIntoWall(width, height int, columns []image.Image) image.Image {
	bounds := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: width, Y: height},
	}

	wall := image.NewRGBA(bounds)

	wallX := (width - TotalWidth(columns)) / 2
	for _, column := range columns {
		targetRect := image.Rect(
			wallX,
			0,
			wallX+column.Bounds().Dx(),
			column.Bounds().Dy())
		draw.Draw(
			wall,
			targetRect,
			column,
			image.ZP,
			draw.Src)
		wallX += column.Bounds().Dx()
	}

	return wall
}

func SaveWall(wall image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	jpeg.Encode(writer, wall, nil)

	return nil
}
