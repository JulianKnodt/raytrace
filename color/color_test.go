package color

import (
	"image/color"
	"raytrace/utils"
	"testing"
)

func TestConversion(t *testing.T) {
	r, g, b, _ := FromColor(color.Black).ToImageColor().RGBA()
	if r != 0 || g != 0 || b != 0 {
		t.Fail()
	}
}

func TestConversionImage(t *testing.T) {
	img, err := utils.LoadImage("./testdata/dog.jpg")
	if err != nil {
		t.Error(err)
	}

	bounds := img.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			c := img.At(x, y)
			converted := FromColor(c).ToImageColor()
			original := color.RGBAModel.Convert(c)
			if converted != original {
				t.Fail()
			}
		}
	}
}
