package main

import (
    "math"
)
/*
 Intersects: an object should have the 
 ability to check whether it intersects with
 a vector

 Normal: an object should be able to determine a normal on its surface

 Color: an object should have a color
*/
type Object interface {
  Intersects(origin, dir Vec3) (float64, bool)
  Normal(to Vec3) (dir Vec3, invAble bool)
  // invable relates to whether or not the normal can be flipped or not
  // any 2d shape should be invertible
  Color() Vec3
}

type Sphere struct {
  center Vec3
  radiusSqr float64
  color Vec3
}

func NewSphere(center Vec3, radius float64, color Vec3) *Sphere {
  return &Sphere{center, radius * radius, color}
}

func (s Sphere) Normal(p Vec3) (Vec3, bool) {
  return Sub(p, s.center), false
}

func (s Sphere) Color() Vec3 {
  return s.color
}

func (s Sphere) Intersects(origin, dir Vec3) (a float64, didInter bool) {
  center := Sub(s.center, origin)
  toNormal := Dot(center, dir)
  if toNormal < 0 {
    return a, false
  }
  distSqr := Dot(center, center) - toNormal * toNormal
  if distSqr > s.radiusSqr {
    return a, false
  }

  interDist := math.Sqrt(s.radiusSqr - distSqr)
  t0 := toNormal - interDist
  t1 := toNormal + interDist

  if t0 < 0 {
    return t1, true
  } else {
    return t0, true
  }
}

