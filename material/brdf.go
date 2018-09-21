package material

import (
	"math"

	v "github.com/julianknodt/vector"
)

// Ref:
// https://www.cs.cmu.edu/afs/cs/academic/class/15462-f09/www/lec/lec8.pdf

type Scatter interface {
	BidirectionalRadianceDistribution(normal, incident, viewing v.Vec3) float64
}

type BRDF func(n, i, v v.Vec3) float64

func (b BRDF) BidirectionalRadianceDistribution(n, i, v v.Vec3) float64 {
	return b(n, i, v)
}

// Model for lambertian reflectance
// Which reflects equally in all directions
type Lambertian struct {

	// [0,1] Diffuse Reflection
	Albedo float64
}

func (l Lambertian) BidirectionalRadianceDistribution(n, i, v v.Vec3) float64 {
	// math.Cos(ti) should be replaced by incident_direction • normal
	return n.Dot(i) * l.Albedo / math.Pi
}
