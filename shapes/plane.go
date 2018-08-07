package shapes

import (
	m "raytrace/material"
	obj "raytrace/object"
	v "raytrace/vector"
)

type Plane struct {
	point v.Vec3
	norm  v.Vec3
	m.Material
}

// should be open to other constructions
func NewPlane(p, norm v.Vec3, mat m.Material) *Plane {
	return &Plane{p, v.Unit(norm), mat}
}

func (p Plane) Intersects(origin, dir v.Vec3) (float64, obj.SurfaceElement) {
	denom := v.Dot(dir, p.norm)
	if denom == 0 {
		return -1, nil
	}
	param := v.Dot(v.Sub(p.point, origin), p.norm) / denom
	if param < 0 {
		return param, nil
	}
	return param, p
}

func (p Plane) Normal(_to v.Vec3) (v.Vec3, bool) {
	return p.norm, true
}

func (p Plane) Mat() m.Material {
	return p.Material
}
