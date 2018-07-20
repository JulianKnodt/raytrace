package main

import (
	obj "github.com/julianknodt/raytrace/object"
	v "github.com/julianknodt/raytrace/vector"
)

type Plane struct {
	point v.Vec3
	norm  v.Vec3
	color v.Vec3
}

// should be open to other constructions
func NewPlane(p, norm v.Vec3, c v.Vec3) *Plane {
	return &Plane{p, v.Unit(norm), c}
}

func (p Plane) Intersects(origin, dir v.Vec3) (float64, obj.Shape) {
	param := v.Dot(v.Sub(p.point, origin), p.norm) / v.Dot(dir, p.norm)
	if param >= 0 {
		return param, p
	}
	return param, nil
}

func (p Plane) Normal(_to v.Vec3) (v.Vec3, bool) {
	return p.norm, true
}

func (p Plane) Color() v.Vec3 {
	return p.color
}
