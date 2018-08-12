package material

import (
	"image"
)

type ImageMaterial struct {
	image.Image
	currX int
	currY int
}

func NewImageMaterial(img image.Image) *ImageMaterial {
	return &ImageMaterial{
		img,
		0,
		0,
	}
}

func (i *ImageMaterial) Emitted() [3]float64 {
	r, g, b, _ := i.At(i.currX, i.currY).RGBA()
	i.currX++
	if i.currX == i.Bounds().Dx() {
		i.currX = 0
		i.currY++
	}
	if i.currY == i.Bounds().Dy() {
		i.currY = 0
	}
	return [3]float64{float64(uint8(r)), float64(uint8(g)), float64(uint8(b))}
}
