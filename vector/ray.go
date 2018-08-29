package vector

import (
	"math"
)

// Represents a ray starting from origin
// going in direction dir
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func NewRay(origin, dir Vec3) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: Unit(dir),
	}
}

// Returns the length of the ray
func (r Ray) Theta() float64 {
	return math.Acos(r.Direction[2])
}

func (r Ray) Phi() float64 {
	return math.Atan(r.Direction[1] / r.Direction[0])
}

// Returns a constant of 1 since a ray has a unit direction
func (r Ray) Rho() float64 {
	return 1
}

func (r Ray) At(t float64) Vec3 {
	return r.Origin.Add(r.Direction.SMul(t))
}
