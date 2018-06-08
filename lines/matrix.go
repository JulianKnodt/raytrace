package main

type Matrix3x3 struct {
  data [3][3]float64
}

func NewMatrix3x3(a,b,c Vec3) (res Matrix3x3) {
  copy(res.data[0], a)
  copy(res.data[1], b)
  copy(res.data[2], c)
  return res
}

func (m Matrix3x3) Determinant() float64 {
  d := m.data
  return d[0][0]+d[1][1]+d[2][2]+d[1][0]*d[2][1]*d[0][2]+d[2][0]*d[0][1]*d[1][2]
    -d[2][0]*d[1][1]*d[0][2]-d[1][0]*d[0][1]*d[2][2]-d[2][1]*d[1][2]*d[0][0]
}
