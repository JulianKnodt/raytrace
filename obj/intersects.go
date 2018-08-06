package obj

import (
	"math"
	mat "raytrace/material"
	obj "raytrace/object"
	"raytrace/shapes"
	v "raytrace/vector"
)

func (o Obj) FaceN(n int) []v.Vec3 {
	face := o.F[n]
	out := make([]v.Vec3, 0, len(face))
	for _, p := range face {
		switch len(p) {
		case 1:
			single := o.V[p[0]]
			out = append(out, v.Vec3{single[0], single[1], single[2]})
		case 2:
			// TODO
		case 3:
			// TODO
		}
	}
	return out
}

func (o Obj) Intersects(origin, dir v.Vec3) (float64, obj.Shape) {
	min := math.Inf(1)
	var shape obj.Shape
	for i := 0; i < len(o.F); i++ {
		face := o.FaceN(i)
		if t, intersects := v.Intersects(face, origin, dir); intersects && t < min {
			min = t
			shape = shapes.NewTriangle(face[0], face[1], face[2], mat.Placeholder{})
		}
	}
	return min, shape
}
