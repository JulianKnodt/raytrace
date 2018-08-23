package shapes

import (
	"fmt"
	m "raytrace/material"
	v "raytrace/vector"
	"testing"
)

func TestPlaneIntersect(t *testing.T) {
	plane := NewPlane(v.Origin, v.Vec3{0, 1, 0}, &m.BasicBlack)
	param, hit := plane.Intersects(*v.NewRay(v.Vec3{0, 2, 0}, v.Vec3{0, -1, 0}))
	if hit == nil {
		fmt.Println("Plane failed to intersect correctly", param, hit)
		t.FailNow()
	}
}
