package main

type Triangle struct {
  a Vec3
  b Vec3
  c Vec3
  color Vec3
}

func (t Triangle) Intersects(origin, dir Vec3) (float64, bool) {
  edge1 := Sub(t.b, t.a)
  edge2 := Sub(t.c, t.a)
  h := Cross(dir, edge2)
  area := Dot(edge1, h)
  if area > -epsilon && area < epsilon {
    return -1, false // this is collinear
  }

  invArea := 1/area
  s := Sub(origin, t.a)
  u := invArea * Dot(s, h)
  if u < 0 || u > 1 {
    return -1, false
  }

  q := Cross(s, edge1)
  v := invArea * Dot(dir, q)
  if v < 0 || (u + v) > 1 {
    return -1, false
  }

  par := invArea * Dot(edge2, q)
  return par, par > epsilon
}

func (t Triangle) Color() Vec3 {
  return t.color
}

func (t Triangle) Normal(_to Vec3) (Vec3, bool) {
  return Unit(Cross(Sub(t.a, t.b), Sub(t.c, t.a))), true
}

func NewTriangle(a, b, c, color Vec3) *Triangle {
  return &Triangle{a,b,c,color}
}
