package octree

import (
	"math"
	"math/rand"
	"testing"
	"time"

	v "github.com/julianknodt/vector"
	"raytrace/bounding"
	"raytrace/shapes"
)

func RandomItem() OctreeItem {
	return shapes.NewTriangle(*v.RandomVector(), *v.RandomVector(), *v.RandomVector(), nil)
}

func TestOctree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	o := NewEmptyOctree(*bounding.NewBox(0, 1000, 0, 1000, 0, 1000))
	for i := 0; i < 100; i++ {
		o.Insert(RandomItem())
	}
	o.Flatten()

	// shouldn't error or loop infinitely
}

func TestOctreeIntersects(t *testing.T) {
	o := NewEmptyOctree(*bounding.NewBox(-10, 10, -10, 10, -10, 10))
	for i := 0; i < 100; i++ {
		o.Insert(RandomItem())
	}

	o.Insert(shapes.NewTriangle(v.Vec3{0, 1, 0}, v.Vec3{0, 0, 1}, v.Vec3{1, 0, 0}, nil))
	o.Flatten()

	origin := v.Vec3{0, 0, 0}
	dir := v.Vec3{1, 1, 1}

	min, el := o.Intersects(*v.NewRay(origin, dir))
	if math.IsInf(min, 1) || el == nil {
		t.Fail()
	}
}

func BenchmarkOctreeIntersects(b *testing.B) {

}
