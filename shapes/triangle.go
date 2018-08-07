package shapes

import (
	m "raytrace/material"
	obj "raytrace/object"
	vec "raytrace/vector"
)

type Triangle struct {
	a vec.Vec3
	b vec.Vec3
	c vec.Vec3
	m.Material
}

func (t Triangle) Intersects(origin, dir vec.Vec3) (float64, obj.SurfaceElement) {
	if param, intersects := vec.IntersectsTriangle(t.a, t.b, t.c, origin, dir); intersects {
		return param, t
	}
	return -1, nil
}

func (t Triangle) Mat() m.Material {
	return t.Material
}

func (t Triangle) Normal(_to vec.Vec3) (vec.Vec3, bool) {
	return vec.Unit(vec.Cross(vec.Sub(t.a, t.b), vec.Sub(t.c, t.a))), true
}

func NewTriangle(a, b, c vec.Vec3, mat m.Material) *Triangle {
	return &Triangle{a, b, c, mat}
}
