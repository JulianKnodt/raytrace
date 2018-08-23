package shapes

import (
	"fmt"
	m "raytrace/material"
	v "raytrace/vector"
	"testing"
)

func TestSphereIntersect(t *testing.T) {
	sphere := NewSphere(v.Vec3{0, 0, -10}, 3.0, &m.BasicBlack)
	param, hit := sphere.Intersects(*v.NewRay(v.Origin, v.Vec3{0, 0, -1}))
	if hit == nil {
		fmt.Println("Failed sphere intersection", param, hit)
		t.FailNow()
	}
}

func TestSphereIntersect2(t *testing.T) {
	sphere := NewSphere(v.Vec3{0, 0, -10}, 3.0, &m.BasicBlack)
	param, hit := sphere.Intersects2(*v.NewRay(v.Origin, v.Vec3{0, 0, -1}))
	if hit == nil {
		fmt.Println("Failed sphere intersection2", param, hit)
		t.FailNow()
	}
}
