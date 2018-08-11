package mesh

import (
	"math"
	"raytrace/object"
	v "raytrace/vector"
)

type LazyMesh struct {
	NumVertices int
	Face        func(int) []v.Vec3
	FaceTexture func(int) object.SurfaceElement
}

func (lm LazyMesh) Intersects(origin, dir v.Vec3) (min float64, shape object.SurfaceElement) {
	min = math.Inf(1)
	shape = nil
	for i := 0; i < lm.NumVertices; i++ {
		face := lm.Face(i)
		if t, intersects := v.Intersects(face, origin, dir); intersects && t < min {
			min = t
			shape = lm.FaceTexture(i)
		}
	}
	return
}
