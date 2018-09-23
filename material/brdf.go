package material

import (
	v "github.com/julianknodt/vector"
)

// Ref:
// https://www.cs.cmu.edu/afs/cs/academic/class/15462-f09/www/lec/lec8.pdf

type Scatter interface {
	BidirectionalRadianceDistribution(normal, incident, viewing v.Ray) float64
}

// n is normal, i is incident, c is camera(or viewer)
type BRDF func(n, i, c v.Ray) float64

func (b BRDF) BidirectionalRadianceDistribution(n, i, c v.Ray) float64 {
	return b(n, i, c)
}

// Model for lambertian reflectance
// Which reflects equally in all directions
type Lambertian struct {

	// [0,1] Diffuse Reflection
	// http://www.curry.eas.gatech.edu/Courses/6140/ency/Chapter9/Ency_Atmos/Reflectance_Albedo_Surface.pdf
	Albedo float64
}

func (l Lambertian) BidirectionalRadianceDistribution(n, i, c v.Ray) float64 {
	// math.Cos(ti) should be replaced by incident_direction • normal
	return c.Direction.Dot(i.Direction) * l.Albedo
}
