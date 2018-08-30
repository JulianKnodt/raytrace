package color

import (
	"image/color"
	v "raytrace/vector"
)

// A representation of a color w/ RGB values
// in [0, 1]
// And alpha as represented as uint8
type Normalized struct {
	// RGB is the RGB of the color
	RGB v.Vec3
	// Alpha is still only a uint32
	A uint8
}

const maxUint32 = 0xFFFF
const maxUint8 = 0xFF

func FromColor(c color.Color) Normalized {
	r, g, b, a := c.RGBA()
	return Normalized{
		v.Vec3{
			float64(r) / maxUint32, float64(g) / maxUint32, float64(b) / maxUint32,
		},
		uint8(a / maxUint32),
	}
}

func FromNormalized(a, b, c, max float64) Normalized {
	return Normalized{
		v.Vec3{
			a / max, b / max, c / max,
		},
		maxUint8,
	}
}

func (n Normalized) Uint8() v.Vec3 {
	return n.RGB.SMul(maxUint8)
}

func (r Normalized) ToImageColor() color.RGBA {
	scaled := r.Uint8()
	return color.RGBA{
		R: uint8(scaled[0]),
		G: uint8(scaled[1]),
		B: uint8(scaled[2]),
		A: uint8(r.A),
	}
}
