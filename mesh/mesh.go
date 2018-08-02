package mesh

import (
	"math"
	mat "raytrace/material"
	obj "raytrace/object"
	"raytrace/shapes"
	v "raytrace/vector"
)

type Mesh struct {
	numVertices uint64
	numFaces    uint64
	Vertices    []v.Vec3
	Order       [][]int
}

func (m Mesh) FaceN(n uint64) []v.Vec3 {
	order := m.Order[n]
	out := make([]v.Vec3, 0, len(order))
	for _, v := range order {
		out = append(out, m.Vertices[v])
	}
	return out
}

func (m Mesh) Verts() uint64 {
	return m.numVertices
}

func (m Mesh) Faces() uint64 {
	return m.numFaces
}

func (m Mesh) Edges() (count uint64) {
	for i := uint64(0); i < m.Faces(); i++ {
		count += uint64(len(m.FaceN(i)))
	}
	return
}

func (m Mesh) Intersects(origin, dir v.Vec3) (float64, obj.Shape) {
	min := math.Inf(1)
	var shape obj.Shape
	for i := uint64(0); i < m.Faces(); i++ {
		face := m.FaceN(i)
		if t, intersects := v.Intersects(face, origin, dir); intersects && t < min {
			min = t
			shape = shapes.NewTriangle(face[0], face[1], face[2], mat.Placeholder{})
		}
	}
	return min, shape
}
