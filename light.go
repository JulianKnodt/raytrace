package main

/*
  A light should be able to determine whether
  or not it can shine a light in a certain direction
  it should be noted that in the case of raytracing,
  the light goes from the point into the light

  A light should have a color
*/
type Light interface {
  LightTo(point Vec3) (dir Vec3, canIllum bool)
  Color() Vec3
}

// light which emits light in all directions
type PointLight struct {
  center Vec3
  color Vec3
}


func (p PointLight) Color() Vec3 {
  return p.color
}

func (p PointLight) LightTo(point Vec3) (dir Vec3, canIllum bool) {
  return Unit(Sub(p.center, point)), true
}
