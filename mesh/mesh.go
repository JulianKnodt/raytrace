package mesh

import (
	"math"
	mat "raytrace/material"
	obj "raytrace/object"
	"raytrace/shapes"
	v "raytrace/vector"
)

type Mesh struct {
	numVertices int
	Vertices    []v.Vec3
	Order       [][]int
	Materials   []mat.Material
}

func (m Mesh) FaceN(n int) []v.Vec3 {
	order := m.Order[n]
	out := make([]v.Vec3, 0, len(order))
	for _, p := range order {
		out = append(out, m.Vertices[p])
	}
	return out
}

func (m Mesh) MaterialN(n int) mat.Material {
	var material mat.Material = mat.Placeholder{}
	if len(m.Materials) > n {
		material = m.Materials[n]
	}
	return material
}

func (m Mesh) Verts() int {
	return m.numVertices
}

func (m Mesh) Faces() int {
	return len(m.Order)
}

func (m Mesh) Edges() (count int) {
	for i := 0; i < m.Faces(); i++ {
		count += len(m.FaceN(i))
	}
	return
}

func (m Mesh) Intersects(r v.Ray) (float64, obj.SurfaceElement) {
	min := math.Inf(1)
	var shape obj.SurfaceElement
	for i := 0; i < m.Faces(); i++ {
		face := m.FaceN(i)
		if t, intersects := v.Intersects(face, r.Origin, r.Direction); intersects && t < min {
			min = t
			shape = shapes.NewTriangle(face[0], face[1], face[2], m.MaterialN(i))
		}
	}
	return min, shape
}
