package sky

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type Sky struct {
	image.Image
}

func NewSky(r io.Reader) (*Sky, error) {
	img, _, err := image.Decode(r)
	return &Sky{
		img,
	}, err
}

func (s Sky) At(x, y int) color.Color {
	if s.Image != nil {
		return s.Image.At(x, y)
	}
	return color.Transparent
}
