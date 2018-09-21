package material

import (
	"image"
	"image/color"
	"math"
	c "raytrace/color"
)

// https://learnopengl.com/Getting-started/Textures
// https://cglearn.codelight.eu/pub/computer-graphics/textures-and-sampling
// http://ogldev.atspace.co.uk/www/tutorial16/tutorial16.html
type Texture interface {
	At(x, y int) color.Color
	Bounds() image.Rectangle
}

// Represents values which u and v can be between
// have to have an epsilon
const maxNormalized = 1
const minNormalized = 0

// coerces x between min and max
func coerceInto(x, min, max float64) float64 {
	return math.Max(min, math.Min(max, x))
}

// u and v represent interpolation between two points
// And are expected to be [0,1]
func Sample(t Texture, u, v float64) c.Normalized {
	u = coerceInto(u, minNormalized, maxNormalized)
	v = coerceInto(v, minNormalized, maxNormalized)

	bounds := t.Bounds()
	x := int(float64(bounds.Dx()) * u)
	y := int(float64(bounds.Dy()) * v)

	return *c.FromColor(t.At(x, y))
}
