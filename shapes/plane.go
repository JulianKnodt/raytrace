package shapes

import (
	m "raytrace/material"
	obj "raytrace/object"
	v "raytrace/vector"
)

type Plane struct {
	point v.Vec3
	norm  v.Vec3
	*m.Material
}

// should be open to other constructions
func NewPlane(p, norm v.Vec3, mat *m.Material) *Plane {
	return &Plane{p, *norm.Unit(), mat}
}

func (p Plane) Intersects(r v.Ray) (float64, obj.SurfaceElement) {
	denom := v.Dot(r.Direction, p.norm)
	if denom == 0 {
		return -1, nil
	}
	param := p.point.Sub(r.Origin).Dot(p.norm) / denom
	if param < 0 {
		return param, nil
	}
	return param, p
}

func (p Plane) NormalAt(v.Vec3) (v.Vec3, bool) {
	return p.norm, false
}

func (p Plane) MaterialAt(v.Vec3) *m.Material {
	return p.Material
}
