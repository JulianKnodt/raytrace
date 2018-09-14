package object

import (
	v "github.com/julianknodt/vector"
	m "raytrace/material"
)

type SurfaceElement interface {
	// 2d shapes can be invertible? It really depends
	// This normal is constant based on wherever it originally intersected
	NormalAt(v.Vec3) (dir v.Vec3, invertible bool)

	MaterialAt(v.Vec3) *m.Material
}
