package main

import (
	v "raytrace/vector"
)

var DefaultCameraDir = v.Vec3{0, 0, -1}

type Camera interface {
	Location() v.Vec3
	Direction() v.Vec3
	FOV() float64
}

type StCamera struct {
	loc v.Vec3
	dir v.Vec3
	fov float64
}

func (s StCamera) Location() v.Vec3 {
	return s.loc
}

func (s StCamera) Direction() v.Vec3 {
	return s.dir
}

func (s StCamera) FOV() float64 {
	return s.fov
}

func NewStCamera(loc, dir v.Vec3, fov float64) StCamera {
	return StCamera{loc, dir, fov}
}
