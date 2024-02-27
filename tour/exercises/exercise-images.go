package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	height, width int
	c             uint8
}

// ColorModel returns the Image's color model.
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (img *Image) At(x, y int) color.Color {
	return color.RGBA{img.c - uint8(x+(2*y)), img.c + uint8(y+x), 200, 255}
}

func main() {
	m := Image{100, 200, 100}
	pic.ShowImage(&m)
}
