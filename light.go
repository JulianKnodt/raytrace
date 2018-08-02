package main

import (
	v "raytrace/vector"
)

/*
  A light should be able to determine whether
  or not it can shine a light in a certain direction
  it should be noted that in the case of raytracing,
  the light goes from the point into the light

  A light should have a color
*/
type Light interface {
	LightTo(point v.Vec3) (dir v.Vec3, canIllum bool)
	Color() v.Vec3
}

// light which emits light in all directions
type PointLight struct {
	center v.Vec3
	color  v.Vec3
}

func (p PointLight) Color() v.Vec3 {
	return p.color
}

func (p PointLight) LightTo(point v.Vec3) (dir v.Vec3, canIllum bool) {
	return v.Unit(v.Sub(p.center, point)), true
}
