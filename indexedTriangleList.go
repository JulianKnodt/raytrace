package main

import (
	"math"
	m "raytrace/material"
	obj "raytrace/object"
	"raytrace/shapes"
	v "raytrace/vector"
)

type IndexedTriangleList struct {
	vertices []v.Vec3
	order    []int
	color    v.Vec3
}

func (t IndexedTriangleList) GetTriangle(nth int) *shapes.Triangle {
	return shapes.NewTriangle(
		t.vertices[t.order[3*nth]],
		t.vertices[t.order[3*nth+1]],
		t.vertices[t.order[3*nth+2]],
		m.Placeholder{})
}

func (t IndexedTriangleList) Size() int {
	return len(t.order) / 3
}

func NewIndexedTriangleList(color v.Vec3) IndexedTriangleList {
	return IndexedTriangleList{make([]v.Vec3, 0), make([]int, 0), color}
}

func (t *IndexedTriangleList) AddPolySet(points []v.Vec3, order []int) {
	t.vertices = append(t.vertices, points...)
	t.order = append(t.order, order...)
}

func (t IndexedTriangleList) Color() v.Vec3 {
	return t.color
}

func (t IndexedTriangleList) Intersects(r v.Ray) (float64, obj.SurfaceElement) {
	var pMax float64 = math.MaxFloat64
	var surface shapes.Triangle
	size := t.Size()
	for i := 0; i < size; i++ {
		curr := t.GetTriangle(i)
		if p, hit := curr.Intersects(r); hit != nil {
			if p < pMax {
				pMax = p
				surface = *curr
			}
		}
	}
	return pMax, surface
}
