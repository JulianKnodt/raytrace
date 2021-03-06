package color

import (
	v "github.com/julianknodt/vector"
	"image/color"
)

// A representation of a color w/ RGB values
// in [0, 1]
// And alpha as represented as uint8
type Normalized struct {
	// Vec is the RGB of the color
	RGB v.Vec3

	A float64
}

const maxUint8 = 0xFF
const maxUint16 = 0xFFFF
const maxUint32 = 0xFFFFFFFF

func FromColor(c color.Color) *Normalized {
	rgba := color.RGBA64Model.Convert(c).(color.RGBA64)
	return &Normalized{
		v.Vec3{
			float64(rgba.R) / maxUint16, float64(rgba.G) / maxUint16, float64(rgba.B) / maxUint16,
		},
		float64(rgba.A) / maxUint16,
	}
}

func FromNormalized(a, b, c, max float64) Normalized {
	return Normalized{
		v.Vec3{
			a / max, b / max, c / max,
		},
		1,
	}
}

func (n Normalized) Uint8() v.Vec3 {
	return *n.RGB.SMul(maxUint8)
}

func (n *Normalized) ToImageColor() color.RGBA {
	if n == nil {
		return color.RGBA{0, 0, 0, 0}
	}
	scaled := n.Uint8()
	out := color.RGBA{
		R: uint8(scaled[0]),
		G: uint8(scaled[1]),
		B: uint8(scaled[2]),
		A: uint8(n.A * maxUint8),
	}
	return out
}

func (n *Normalized) Mix(o Normalized) *Normalized {
	n.RGB = *n.RGB.Add(o.RGB).SMulSet(0.5)
	n.A = (n.A + o.A) / 2
	return n
}

func (n *Normalized) MixVec(vec v.Vec3) *Normalized {
	n.RGB = *n.RGB.Add(vec).SMulSet(0.5)
	return n
}
