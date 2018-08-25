package vector

// Some testing utils for vectors

import (
	"math/rand"
)

func RandomVector() Vec3 {
	return Vec3{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

func RandomTriple() [3]Vec3 {
	return [3]Vec3{
		RandomVector(),
		RandomVector(),
		RandomVector(),
	}
}
