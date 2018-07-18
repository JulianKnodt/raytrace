package main

import (
	v "github.com/julianknodt/raytrace/vector"
	"math"
)

/*
 Intersects: an object should have the
 ability to check whether it intersects with
 a vector

 Normal: an object should be able to determine a normal on its surface

 Color: an object should have a color
*/
type Object interface {
	Intersects(origin, dir v.Vec3) (float64, Object)
	Normal(to v.Vec3) (dir v.Vec3, invAble bool)
	// invable relates to whether or not the normal can be flipped or not
	// any 2d shape should be invertible
	Color() v.Vec3
}

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

func (s Sphere) Intersects(origin, dir v.Vec3) (a float64, intersectingObject Object) {
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
