package camera

import (
	"math/rand"
	"testing"
)

// This is a fairly generic test
func TestCamera(t *testing.T) {
	d := DefaultCamera()
	mid := d.RayTo(0.5, 0.5)
	if mid.Direction != d.Transform.Direction {
		t.Errorf("Incorrect direction for RayTo expected %v, got %v",
			d.Transform.Direction,
			mid.Direction,
		)
	}
}

func BenchmarkRayTo(b *testing.B) {
	d := DefaultCamera()
	x, y := rand.Float64(), rand.Float64()
	for i := 0; i < b.N; i++ {
		d.RayTo(x, y)
	}
}
