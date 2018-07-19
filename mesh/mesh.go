package mesh

import (
	v "github.com/julianknodt/raytrace/vector"
)

type Mesh struct {
	numVertices uint64
	numFaces    uint64
	Vertices    []v.Vec3
	Order       [][]int
}

func (m Mesh) FaceN(n int) []v.Vec3 {
	order := m.Order[n]
	out := make([]v.Vec3, 0, len(order))
	for _, v := range order {
		out = append(out, m.Vertices[v])
	}
	return out
}
