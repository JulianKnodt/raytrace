package shapes

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMaxMin(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	x, y, z := rand.Float64(), rand.Float64(), rand.Float64()

	for i := 0; i < b.N; i++ {
		maxmin(x, y, z)
	}
}
