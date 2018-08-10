package octree

import (
	"testing"
)

func TestSphereIntersects(t *testing.T) {
	s1 := BoundingSphere{
		Center: [3]float64{0, 0, 0},
		Radius: 5,
	}
	s2 := BoundingSphere{
		Center: [3]float64{0, 0, 0},
		Radius: 5,
	}

	if !s1.Intersects(s2) {
		t.Fail()
	}
}

func TestNaiveBoxIntersects(t *testing.T) {
	s1 := NaiveBoundingBox{
		0, 2, 0, 2, 0, 2,
	}
	s2 := NaiveBoundingBox{
		0, 2, 0, 2, 0, 2,
	}

	if !s1.Intersects(s2) {
		t.Fail()
	}
}

func TestNaiveBoxNotIntersects(t *testing.T) {
	s1 := NaiveBoundingBox{
		0, 2, 0, 2, 0, 2,
	}
	s2 := NaiveBoundingBox{
		3, 5, 3, 5, 3, 5,
	}

	if s1.Intersects(s2) {
		t.Fail()
	}
}

func TestNaiveBoxContains(t *testing.T) {
	s1 := NaiveBoundingBox{
		0, 4, 0, 4, 0, 4,
	}
	s2 := NaiveBoundingBox{
		1, 3, 1, 3, 1, 3,
	}

	if !s1.Contains(s2) {
		t.Fail()
	}
}
