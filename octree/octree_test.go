package octree

import (
	"math/rand"
	"testing"
	"time"

	"raytrace/bounding"
	"raytrace/shapes"
	v "raytrace/vector"
)

func RandomVec() v.Vec3 {
	return v.Vec3{
		rand.Float64()*999 - 999,
		rand.Float64()*999 - 999,
		rand.Float64()*999 - 999,
	}
}

func RandomItem() OctreeItem {
	return shapes.NewTriangle(RandomVec(), RandomVec(), RandomVec(), nil)
}

func TestOctree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	o := NewEmptyOctree(*bounding.NewOriginAABB(1000.0))
	for i := 0; i < 100; i++ {
		o.Insert(RandomItem())
	}
	o.Flatten()

	// shouldn't error lol
}
