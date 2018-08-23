package obj

import (
	"math"
	mat "raytrace/material"
	obj "raytrace/object"
	"raytrace/shapes"
	v "raytrace/vector"
)

// Only returns the triangle, meaning it excludes all material related things
func (o Obj) ShapeN(n int) []v.Vec3 {
	face := o.F[n]
	out := make([]v.Vec3, 0, len(face.Elements))
	for _, p := range face.Elements {
		out = append(out, v.Vec3(o.V[p.V]))
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

func (o Obj) Intersects(r v.Ray) (float64, obj.SurfaceElement) {
	min := math.Inf(1)
	var shape obj.SurfaceElement
	for i := 0; i < len(o.F); i++ {
		face := o.ShapeN(i)
		if t, intersects := v.Intersects(face, r.Origin, r.Direction); intersects && t < min {
			min = t
			shape = shapes.NewTriangle(face[0], face[1], face[2], mat.Placeholder{})
		}
	}
	return min, shape
}
