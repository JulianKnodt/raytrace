package main

import (
	vec "github.com/julianknodt/raytrace/vector"
)

type Triangle struct {
	a     vec.Vec3
	b     vec.Vec3
	c     vec.Vec3
	color vec.Vec3
}

func (t Triangle) Intersects(origin, dir vec.Vec3) (float64, Object) {
	edge1 := vec.Sub(t.b, t.a)
	edge2 := vec.Sub(t.c, t.a)
	h := vec.Cross(dir, edge2)
	area := vec.Dot(edge1, h)
	if area > -epsilon && area < epsilon {
		return -1, nil // this is collinear
	}

	invArea := 1 / area
	s := vec.Sub(origin, t.a)
	u := invArea * vec.Dot(s, h)
	if u < 0 || u > 1 {
		return -1, nil
	}

	q := vec.Cross(s, edge1)
	v := invArea * vec.Dot(dir, q)
	if v < 0 || (u+v) > 1 {
		return -1, nil
	}

	par := invArea * vec.Dot(edge2, q)
	if par > epsilon {
		return par, t
	}
	return par, nil
}

func (t Triangle) Color() vec.Vec3 {
	return t.color
}

func (t Triangle) Normal(_to vec.Vec3) (vec.Vec3, bool) {
	return vec.Unit(vec.Cross(vec.Sub(t.a, t.b), vec.Sub(t.c, t.a))), true
}

func NewTriangle(a, b, c, color vec.Vec3) *Triangle {
	return &Triangle{a, b, c, color}
}
