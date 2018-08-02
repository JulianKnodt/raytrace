package shapes

import (
	"fmt"
	v "raytrace/vector"
	"testing"
)

func TestIntersect(t *testing.T) {
	sphere := NewSphere(v.Vec3{0, 0, -10}, 3.0, v.Vec3{0, 0, 0})
	param, hit := sphere.Intersects(v.Origin, v.Vec3{0, 0, -1})
	if hit == nil {
		fmt.Println("Failed sphere intersection", param, hit)
		t.FailNow()
	}
}

func TestIntersect2(t *testing.T) {
	sphere := NewSphere(v.Vec3{0, 0, -10}, 3.0, v.Vec3{0, 0, 0})
	param, hit := sphere.Intersects2(v.Origin, v.Vec3{0, 0, -1})
	if hit == nil {
		fmt.Println("Failed sphere intersection2", param, hit)
		t.FailNow()
	}
}
