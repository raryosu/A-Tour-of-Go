package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	image_func := func(x, y int) uint8 {
		return uint8(x ^ y)
	}
	v := image_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
