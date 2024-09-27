package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	value := byte(x^y)
	return color.RGBA{value, value, 255, 255}
}

func main() {
	m := Image{
		w: 400,
		h: 400,
	}
	pic.ShowImage(m)
}
