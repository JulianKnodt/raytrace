package main

import (
	obj "github.com/julianknodt/raytrace/object"
	v "github.com/julianknodt/raytrace/vector"
	"math"
)

type Sphere struct {
	center    v.Vec3
	radiusSqr float64
	color     v.Vec3
}

func NewSphere(center v.Vec3, radius float64, color v.Vec3) *Sphere {
	return &Sphere{center, radius * radius, color}
}

func (s Sphere) Normal(p v.Vec3) (v.Vec3, bool) {
	return v.Sub(p, s.center), false
}

func (s Sphere) Color() v.Vec3 {
	return s.color
}

func (s Sphere) Intersects(origin, dir v.Vec3) (a float64, shape obj.Shape) {
	center := v.Sub(s.center, origin)
	toNormal := v.Dot(center, dir)
	if toNormal < 0 {
		return a, nil
	}
	distSqr := v.Dot(center, center) - toNormal*toNormal
	if distSqr > s.radiusSqr {
		return a, nil
	}

	interDist := math.Sqrt(s.radiusSqr - distSqr)
	t0 := toNormal - interDist
	t1 := toNormal + interDist

	if t0 < 0 {
		return t1, s
	} else {
		return t0, s
	}
}
