package main

import (
  "math"
)

type IndexedTriangleList struct {
	vertices []Vec3
  order []int
  color Vec3
}

func (t IndexedTriangleList) GetTriangle(nth int) Triangle {
	return Triangle{
    t.vertices[t.order[3*nth]],
		t.vertices[t.order[3*nth+1]],
		t.vertices[t.order[3*nth+2]],
    t.color}
}

func (t IndexedTriangleList) Size() int {
	return len(t.order) / 3
}

func (t *IndexedTriangleList) AddTriangle(tri Triangle) {
	t.vertices = append(t.vertices, tri.a, tri.b, tri.c)
}

func NewIndexedTriangleList(color Vec3) IndexedTriangleList {
	return IndexedTriangleList{make([]Vec3, 0), make([]int, 0), color}
}

func (t *IndexedTriangleList) AddPolySet(points []Vec3, order []int) {
  t.vertices = append(t.vertices, points...)
  t.order = append(t.order, order...)
}

func (t IndexedTriangleList) Color() Vec3 {
  return t.color
}

func (t IndexedTriangleList) Intersects(origin, dir Vec3) (float64, Object) {
  var pMax float64 = math.MaxFloat64
  var surface Triangle
  size := t.Size()
  for i := 0; i < size; i ++ {
    curr := t.GetTriangle(i)
    if p, hit := curr.Intersects(origin, dir); hit != nil {
      if p < pMax {
        pMax = p
        surface = curr
      }
    }
  }
  return pMax, surface
}
