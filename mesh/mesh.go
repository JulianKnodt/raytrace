package mesh

import (
	obj "github.com/julianknodt/raytrace/object"
	v "github.com/julianknodt/raytrace/vector"
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

func (m Mesh) Intersects(origin, dir v.Vec3) (float64, obj.Shape) {
	for i := uint64(0); i < m.Faces(); i++ {
		v.Intersects(m.FaceN(i), origin, dir)
	}
	return -1, nil
}
