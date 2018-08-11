package material

import (
	"math"
	v "raytrace/vector"
)

type Impulse struct {
	Magnitude [3]float64
	Direction [3]float64
}

func square(t float64) float64 {
	return t * t
}

// Returns the fresnel representation of the color from
// cosTheta where theta is the viewing angle
func Fresnel(color v.Vec3, cosTheta float64) v.Vec3 {
	return v.Lerp(
		color,
		v.Sub(v.Vec3{1, 1, 1}, color),
		math.Pow(1-cosTheta, 5),
	)
}
