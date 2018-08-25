package material

import (
	"image"
	"math"
	"raytrace/color"
)

// https://learnopengl.com/Getting-started/Textures
// https://cglearn.codelight.eu/pub/computer-graphics/textures-and-sampling
// http://ogldev.atspace.co.uk/www/tutorial16/tutorial16.html
type Texture interface {
	// u and v represent interpolation between two points
	// And are expected to be [0,1]
	Sample(u, v float64) color.RGBA
}

type ImageTexture struct {
	image.Image
}

// Represents values which u and v can be between
// have to have an epsilon
const maxNormalized = 1
const minNormalized = 0

func coerceInto(x, min, max float64) float64 {
	return math.Max(min, math.Min(max, x))
}

func (it ImageTexture) Sample(u, v float64) color.RGBA {
	u = coerceInto(u, minNormalized, maxNormalized)
	v = coerceInto(v, minNormalized, maxNormalized)

	bounds := it.Image.Bounds()
	x := int(float64(bounds.Dx()) * u)
	y := int(float64(bounds.Dy()) * v)

	return color.FromColor(it.Image.At(x, y))
}
