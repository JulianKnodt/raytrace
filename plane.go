package main

type Plane struct {
  point Vec3
  norm Vec3
  color Vec3
}

// should be open to other constructions
func NewPlane(p, norm Vec3, c Vec3) *Plane {
  return &Plane{p, Unit(norm), c}
}

func (p Plane) Intersects(origin, dir Vec3) (float64, Object) {
  param := Dot(Sub(p.point, origin), p.norm)/Dot(dir, p.norm)
  if param >= 0 {
    return param, p
  }
  return param, nil
}

func (p Plane) Normal(_to Vec3) (Vec3, bool) {
  return p.norm, true
}

func (p Plane) Color() Vec3 {
  return p.color
}

