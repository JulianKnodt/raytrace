package shapes

import (
	"math"

	"raytrace/bounding"
	m "raytrace/material"
	obj "raytrace/object"
	vec "raytrace/vector"
)

// Represents a triangle with a material and textured vertices
type Triangle struct {
	// Coordinates of vertices within world
	v0, v1, v2 vec.Vec3

	// Texture Coordinates
	t0, t1, t2 vec.Vec3

	// Mapped Normal Coordinates
	n0, n1, n2 vec.Vec3

	// Rendered material of this triangle
	*m.Material
}

func (t Triangle) Intersects(r vec.Ray) (float64, obj.SurfaceElement) {
	if param, intersects := vec.IntersectsTriangle(t.v0, t.v1, t.v2,
		r.Origin, r.Direction); intersects {
		return param, t
	}
	return math.Inf(1), nil
}

func (t Triangle) MaterialAt(vec.Vec3) *m.Material {
	return t.Material
}

func (t Triangle) NormalAt(v vec.Vec3) (vec.Vec3, bool) {
	b := t.Barycentric(v)
	return t.n0.SMul(b[0]).Add(t.n1.SMul(b[1])).Add(t.n2.SMul(b[2])), true
}

func NewTriangle(a, b, c vec.Vec3, mat *m.Material) *Triangle {
	n := a.Sub(c).Cross(a.Sub(b)).Unit()
	return &Triangle{
		v0: a, v1: b, v2: c,
		n0: n, n1: n, n2: n,
		Material: mat,
	}
}

func (t *Triangle) SetTextures(t0, t1, t2 vec.Vec3) {
	t.t0 = t0
	t.t1 = t1
	t.t2 = t2
}

func (t *Triangle) SetNormals(n0, n1, n2 vec.Vec3) {
	t.n0 = n0
	t.n1 = n1
	t.n2 = n2
}

func ToTriangles(vecs []vec.Vec3, mat *m.Material) []Triangle {
	if len(vecs) < 3 {
		return nil
	}
	out := make([]Triangle, 0, len(vecs)-2)

	for i := 0; i < len(vecs)-2; i++ {
		out = append(out, *NewTriangle(vecs[0], vecs[1], vecs[2], mat))
	}

	return out
}

// Returns the bounding box for the triangle
func (t Triangle) Box() bounding.Box {
	return bounding.Box{
		Min: t.v0.Min(t.v1).Min(t.v2),
		Max: t.v0.Max(t.v1).Max(t.v2),
	}
}

func (t Triangle) Area() float64 {
	n := t.v0.Sub(t.v1).Cross(t.v0.Sub(t.v2))
	return n.Magn() / 2
}

// https://gamedev.stackexchange.com/questions/23743/whats-the-most-efficient-way-to-find-barycentric-coordinates
func (t Triangle) Barycentric(v vec.Vec3) vec.Vec3 {
	v0 := t.v0.Sub(t.v1)
	v1 := t.v0.Sub(t.v2)
	v2 := v.Sub(t.v0)
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
