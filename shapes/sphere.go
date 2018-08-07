package shapes

import (
	"math"
	m "raytrace/material"
	obj "raytrace/object"
	v "raytrace/vector"
)

type Sphere struct {
	center    v.Vec3
	radiusSqr float64
	m.Material
}

func NewSphere(center v.Vec3, radius float64, mat m.Material) *Sphere {
	return &Sphere{center, radius * radius, mat}
}

func (s Sphere) Normal(p v.Vec3) (v.Vec3, bool) {
	return v.Sub(p, s.center), false
}

func (s Sphere) Mat() m.Material {
	return s.Material
}

func (s Sphere) Intersects(origin, dir v.Vec3) (a float64, shape obj.SurfaceElement) {
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

func (s Sphere) Intersects2(origin, dir v.Vec3) (t float64, shape obj.SurfaceElement) {
	centerDiff := v.Sub(origin, s.center)
	a := v.SqrMagn(dir)
	b := 2 * v.Dot(dir, centerDiff)
	c := v.SqrMagn(centerDiff) - s.radiusSqr
	discrim := (b * b) - (4 * a * c)
	if discrim < 0 || a == 0 {
		return t, nil
	}

	if v.SqrMagn(centerDiff) <= s.radiusSqr {
		return t, nil
	}

	t0 := (-b + math.Sqrt(discrim)) / (2 * a)
	t1 := (-b - math.Sqrt(discrim)) / (2 * a)
	shape = s
	switch {
	case t0 < 0 && t1 < 0:
		return t, nil
	case t0 < 0:
		t = t1
	case t1 < 0:
		t = t0
	default:
		// return closest point
		t = math.Min(t0, t1)
	}
	return
}
