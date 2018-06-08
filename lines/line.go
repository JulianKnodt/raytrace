package main

type Line struct {
  origin Vec3
  dir Vec3
  thickness float64
  color Vec3
}

func (l Line) Normal(to Vec3) (Vec3, bool) {
  toDir := Sub(to, l.origin)
  return SMul(Dot(to, l.dir)/SqrMagn(toDir), toDir), false
}

//func (l Line) Intersects(origin, dir Vec3) (float64, bool) {
  // take some sort of determinant here 
//  col1 := Add(l.origin, origin)
//  col2 := 
//}
