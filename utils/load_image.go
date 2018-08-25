package utils

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// Load a png file from a path
func LoadImage(filepath string) (image.Image, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	f.Close()
	return img, err
}
