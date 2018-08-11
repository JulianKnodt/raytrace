package object

import (
	m "raytrace/material"
	v "raytrace/vector"
)

type SurfaceElement interface {
	// 2d shapes can be invertible? It really depends
	// This normal is constant based on wherever it originally intersected
	Normal() (dir v.Vec3, invertible bool)

	Mat() m.Material
}

type Surfel struct {
	Norm     v.Vec3
	Material m.Material
}

func (s Surfel) Normal() (v.Vec3, bool) {
	return s.Norm, false
}

func (s Surfel) Mat() m.Material {
	return s.Material
}

type InvertibleSurfel struct {
	Norm     v.Vec3
	Material m.Material
}

func (s InvertibleSurfel) Normal() (v.Vec3, bool) {
	return s.Norm, true
}

func (s InvertibleSurfel) Mat() m.Material {
	return s.Material
}
