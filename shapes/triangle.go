package shapes

import (
	obj "github.com/julianknodt/raytrace/object"
	vec "github.com/julianknodt/raytrace/vector"
)

type Triangle struct {
	a     vec.Vec3
	b     vec.Vec3
	c     vec.Vec3
	color vec.Vec3
}

func (t Triangle) Intersects(origin, dir vec.Vec3) (float64, obj.Shape) {
	if param, intersects := vec.IntersectsTriangle(t.a, t.b, t.c, origin, dir); intersects {
		return param, t
	}
	return -1, nil
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
