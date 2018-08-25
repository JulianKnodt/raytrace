package color

import (
	"image"
	"image/png"
	"os"
)

const swatchSize = 256

// Creates a file of a color
// For testing
func (r RGBA) Swatch(swatchFile string) error {
	f, err := os.Create(swatchFile)
	if err != nil {
		return err
	}
	defer f.Close()

	c := r.ToImageColor()

	img := image.NewRGBA(image.Rect(0, 0, swatchSize, swatchSize))
	for x := 0; x < swatchSize; x++ {
		for y := 0; y < swatchSize; y++ {
			img.SetRGBA(x, y, c)
		}
	}

	return png.Encode(f, img)
}
