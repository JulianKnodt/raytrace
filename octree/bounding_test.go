package octree

import (
	v "raytrace/vector"
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

func TestAABBBoxIntersects(t *testing.T) {
	s1 := AxisAlignedBoundingBox{
		0, 2, 0, 2, 0, 2,
	}
	s2 := AxisAlignedBoundingBox{
		0, 2, 0, 2, 0, 2,
	}

	if !s1.Intersects(s2) {
		t.Fail()
	}
}

func TestAABBBoxNotIntersects(t *testing.T) {
	s1 := AxisAlignedBoundingBox{
		0, 2, 0, 2, 0, 2,
	}
	s2 := AxisAlignedBoundingBox{
		3, 5, 3, 5, 3, 5,
	}

	if s1.Intersects(s2) {
		t.Fail()
	}
}

func TestAABBBoxContains(t *testing.T) {
	s1 := AxisAlignedBoundingBox{
		0, 4, 0, 4, 0, 4,
	}
	s2 := AxisAlignedBoundingBox{
		1, 3, 1, 3, 1, 3,
	}

	if !s1.Contains(s2) {
		t.Fail()
	}
}

func TestAABBIntersectsRay(t *testing.T) {
	s1 := AxisAlignedBoundingBox{
		1, 3, 1, 3, 1, 3,
	}
	origin := v.Vec3{0, 0, 0}
	dir := v.Vec3{1, 1, 1}
	if !s1.IntersectsRay(origin, dir) {
		t.Fail()
	}

	origin = v.Vec3{0, 2, 0}
	dir = v.Vec3{1, 0, 1}
	if !s1.IntersectsRay(origin, dir) {
		t.Fail()
	}
}

func TestAABBIntersectsRayPastBox(t *testing.T) {
	s1 := AxisAlignedBoundingBox{
		1, 3, 1, 3, 1, 3,
	}
	origin := v.Vec3{5, 5, 5}
	dir := v.Vec3{1, 1, 1}
	if s1.IntersectsRay(origin, dir) {
		t.Fail()
	}
}

func BenchmarkAABBIntersectsRay(b *testing.B) {
	s1 := AxisAlignedBoundingBox{
		1, 3, 1, 3, 1, 3,
	}
	origin := v.Vec3{0, 0, 0}
	dir := v.Vec3{1, 1, 1}
	for i := 0; i < b.N; i++ {
		s1.IntersectsRay(origin, dir)
	}
}
