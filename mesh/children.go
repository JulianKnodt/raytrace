package mesh

import (
	"raytrace/octree"
	"raytrace/shapes"
)

func (m Mesh) Children() []octree.OctreeItem {
	numFaces := m.Faces()
	out := make([]octree.OctreeItem, 0, numFaces)
	for i := 0; i < numFaces; i++ {
		face := m.FaceN(i)
		for j := 0; j < len(face)-2; j++ {
			out = append(out,
				shapes.NewTriangle(
					face[j],
					face[j+1],
					face[j+2],
					m.MaterialN(i),
				),
			)
		}
	}
	return out
}
