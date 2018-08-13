package shapes

import (
	"math"
	m "raytrace/material"
	obj "raytrace/object"
	"raytrace/octree"
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

func (t Triangle) Normal() (vec.Vec3, bool) {
	return vec.Unit(vec.Cross(vec.Sub(t.a, t.b), vec.Sub(t.c, t.a))), true
}

func NewTriangle(a, b, c vec.Vec3, mat m.Material) *Triangle {
	return &Triangle{a, b, c, mat}
}

func maxmin(a, b, c float64) (max, min float64) {
	switch {
	case a >= b && a >= c:
		return a, math.Min(b, c)
	case b >= a && b >= c:
		return b, math.Min(a, c)
	case c >= a && c >= b:
		return c, math.Min(b, a)
	default:
		panic("Somehow there is not a min/max amidst three floats")
	}
}

func (t Triangle) Box() octree.AxisAlignedBoundingBox {
	maxX, minX := maxmin(t.a[0], t.b[0], t.c[0])
	maxY, minY := maxmin(t.a[1], t.b[1], t.c[1])
	maxZ, minZ := maxmin(t.a[2], t.b[2], t.c[2])
	return octree.AxisAlignedBoundingBox{
		Xx: minX,
		XX: maxX,
		Yy: minY,
		YY: maxY,
		Zz: minZ,
		ZZ: maxZ,
	}
}
