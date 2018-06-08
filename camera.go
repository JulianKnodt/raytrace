package main

var DefaultCameraDir = Vec3{0,0,-1}

type Camera interface {
  Location() Vec3
  Direction() Vec3
  FOV() float64
}

type StCamera struct {
  loc Vec3
  dir Vec3
  fov float64
}

func (s StCamera) Location() Vec3 {
  return s.loc
}

func (s StCamera) Direction() Vec3 {
  return s.dir
}

func (s StCamera) FOV() float64 {
  return s.fov
}

func NewStCamera(loc, dir Vec3, fov float64) StCamera {
  return StCamera{loc, dir, fov}
}

