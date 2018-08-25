package color

import (
	"image/color"
	v "raytrace/vector"
)

// Differs from go's color package by using floats
// instead of uints
type RGBA struct {
	// RGB is the RGB of the color
	RGB v.Vec3
	// Alpha is still only a uint32
	A uint8
}

func FromColor(c color.Color) RGBA {
	conv := color.RGBAModel.Convert(c).(color.RGBA)
	return RGBA{
		v.Vec3{float64(conv.R), float64(conv.G), float64(conv.B)},
		conv.A,
	}
}

func FromVector(v v.Vec3) RGBA {
	return RGBA{
		v,
		0xFF,
	}
}

func (r RGBA) ToImageColor() color.RGBA {
	return color.RGBA{
		R: uint8(r.RGB[0]),
		G: uint8(r.RGB[1]),
		B: uint8(r.RGB[2]),
		A: uint8(r.A),
	}
}
