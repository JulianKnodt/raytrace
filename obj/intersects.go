package obj

import (
	"math"
	mat "raytrace/material"
	obj "raytrace/object"
	"raytrace/shapes"
	v "raytrace/vector"
)

func convert(a [4]float64) v.Vec3 {
	return v.Vec3{a[0], a[1], a[2]}
}

// Only returns the triangle, meaning it excludes all material related things
func (o Obj) TriangleN(n int) []v.Vec3 {
	face := o.F[n]
	out := make([]v.Vec3, 0, len(face.Elements))
	for _, p := range face.Elements {
		out = append(out, convert(o.V[p.V]))
	}
	return out
}

func (o Obj) TextureN(n int) []v.Vec3 {
	face := o.F[n]
	out := make([]v.Vec3, 0, len(face.Elements))
	for _, p := range face.Elements {
		out = append(out, v.Vec3(o.Vt[p.Vt]))
	}
	return out
}

func (o Obj) Intersects(origin, dir v.Vec3) (float64, obj.SurfaceElement) {
	min := math.Inf(1)
	var shape obj.SurfaceElement
	for i := 0; i < len(o.F); i++ {
		face := o.TriangleN(i)
		if t, intersects := v.Intersects(face, origin, dir); intersects && t < min {
			min = t
			shape = shapes.NewTriangle(face[0], face[1], face[2], mat.Placeholder{})
		}
	}
	return min, shape
}
