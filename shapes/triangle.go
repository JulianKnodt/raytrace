package shapes

import (
	"math"

	"raytrace/bounding"
	m "raytrace/material"
	obj "raytrace/object"
	"raytrace/utils"
	vec "raytrace/vector"
)

type Triangle struct {
	a vec.Vec3
	b vec.Vec3
	c vec.Vec3
	m.Material
}

func (t Triangle) Intersects(r vec.Ray) (float64, obj.SurfaceElement) {
	if param, intersects := vec.IntersectsTriangle(t.a, t.b, t.c,
		r.Origin, r.Direction); intersects {
		return param, t
	}
	return math.Inf(1), nil
}

func (t Triangle) MaterialAt(vec.Vec3) m.Material {
	return t.Material
}

func (t Triangle) NormalAt(vec.Vec3) (vec.Vec3, bool) {
	return vec.Unit(vec.Cross(vec.Sub(t.a, t.b), vec.Sub(t.c, t.a))), true
}

func NewTriangle(a, b, c vec.Vec3, mat m.Material) *Triangle {
	return &Triangle{a, b, c, mat}
}

func ToTriangles(vecs []vec.Vec3, mat m.Material) []Triangle {
	if len(vecs) < 3 {
		return nil
	}
	out := make([]Triangle, 0, len(vecs)-2)

	for i := 0; i < len(vecs)-2; i++ {
		out = append(out, Triangle{vecs[0], vecs[1], vecs[2], mat})
	}

	return out
}

func (t Triangle) Box() bounding.AxisAlignedBoundingBox {
	maxX, minX := utils.Maxmin(t.a[0], t.b[0], t.c[0])
	maxY, minY := utils.Maxmin(t.a[1], t.b[1], t.c[1])
	maxZ, minZ := utils.Maxmin(t.a[2], t.b[2], t.c[2])
	return bounding.AxisAlignedBoundingBox{
		Xx: minX,
		XX: maxX,
		Yy: minY,
		YY: maxY,
		Zz: minZ,
		ZZ: maxZ,
	}
}

func (t Triangle) Area() float64 {
	n, _ := t.NormalAt(vec.Origin)
	return n.Magn() / 2
}

// https://gamedev.stackexchange.com/questions/23743/whats-the-most-efficient-way-to-find-barycentric-coordinates
func (t Triangle) Barycentric(v vec.Vec3) vec.Vec3 {
	v0 := t.a.Sub(t.b)
	v1 := t.a.Sub(t.c)
	v2 := v.Sub(t.a)
	d00 := v0.Dot(v0)
	d01 := v0.Dot(v1)
	d11 := v1.Dot(v1)
	d20 := v2.Dot(v0)
	d21 := v2.Dot(v1)
	denom := d00*d11 - d01*d01
	a := (d11*d20 - d01*d21) / denom
	b := (d00*d21 - d01*d20) / denom
	return vec.Vec3{a, b, 1 - a - b}
}
