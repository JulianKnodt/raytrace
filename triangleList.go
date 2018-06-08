package main

// This is only a partial implementation and is better suited by IndexedTriangleList

type TriangleList struct {
	vertices []Vec3
  color Vec3
}

func (t TriangleList) GetTriangle(nth int) Triangle {
	return Triangle{t.vertices[3*nth],
		t.vertices[3*nth+1],
		t.vertices[3*nth+2],
    t.color}
}

func (t TriangleList) Size() int {
	return len(t.vertices) / 3
}

func (t *TriangleList) AddTriangle(tri Triangle) {
	t.vertices = append(t.vertices, tri.a, tri.b, tri.c)
}

func NewTriangleList(color Vec3) TriangleList {
	return TriangleList{make([]Vec3, 0), color}
}
