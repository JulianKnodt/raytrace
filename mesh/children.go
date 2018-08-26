package mesh

import (
	"raytrace/shapes"
)

func (m Mesh) Children() []shapes.Triangle {
	numFaces := m.Faces()
	out := make([]shapes.Triangle, numFaces)
	for i := 0; i < numFaces; i++ {
	}
	return out
}
