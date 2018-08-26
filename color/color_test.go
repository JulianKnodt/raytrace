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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ApproxEqual(a, b color.RGBA) bool {
	return Abs(int(a.R)-int(b.R)) <= 1 &&
		Abs(int(a.G)-int(b.G)) <= 1 &&
		Abs(int(a.B)-int(b.B)) <= 1 &&
		Abs(int(a.A)-int(b.A)) <= 1
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
			original := color.RGBAModel.Convert(c).(color.RGBA)
			if !ApproxEqual(converted, original) {
				t.Error(converted, original)
			}
		}
	}
}
