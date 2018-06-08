package main

import (
  "math"
)

type IndexedTriangleList struct {
	vertices []Vec3
  order []int
  color Vec3
}

func (t TriangleList) GetTriangle(nth int) Triangle {
	return Triangle{
    t.vertices[t.order[3*nth]],
		t.vertices[t.order[3*nth+1]],
		t.vertices[t.order[3*nth+2]],
    t.color}
}

func (t TriangleList) Size() int {
	return len(t.order) / 3
}

func (t *TriangleList) AddTriangle(tri Triangle) {
	t.vertices = append(t.vertices, tri.a, tri.b, tri.c)
}

func NewTriangleList(color Vec3) TriangleList {
	return TriangleList{make([]Vec3, 0), make([]int, 0), color}
}

func (t *TriangleList) AddPolySet(points []Vec3, order []int) {
  t.vertices = append(t.vertices, points)
  t.order = append(t.order, order)
}

func (t TriangleList) Color() Vec3 {
  return t.color
}

func (t TriangleList) Intersects(origin, dir Vec3) (float64, bool) {
  var pMax float64 = math.MaxFloat64
  var surface Triangle
  size := t.Size()
  for i := 0; i < size; i ++ {
    curr := t.GetTriangle(i)
    if p, hit := curr.Intersects(origin, dir); hit {
      if p < pMax {
        pMax = p
        surface = curr
      }
    }
  }
  return pMax, surface != nil
}
