package octree

import (
	"math"
)

type BoundingSphere struct {
	Center [3]float64
	Radius float64
}

func (b BoundingSphere) Intersects(other BoundingSphere) bool {
	dist := math.Sqrt(
		sqr(b.Center[0]-other.Center[0]) +
			sqr(b.Center[1]-other.Center[1]) +
			sqr(b.Center[2]-other.Center[2]))

	return dist < (b.Radius + other.Radius)
}

func sqr(a float64) float64 {
	return a * a
}

// This is assumed to just be a flat box
type NaiveBoundingBox struct {
	// Min x
	Xx float64
	// Max X
	XX float64
	// Min y
	Yy float64
	// Max y
	YY float64
	// Min z
	Zz float64
	// Max z
	ZZ float64
}

func (n NaiveBoundingBox) Intersects(o NaiveBoundingBox) bool {
	return n.XX > o.Xx && o.XX > n.Xx &&
		n.YY > o.Yy && o.YY > n.Yy &&
		n.ZZ > o.Zz && o.ZZ > n.Zz
}

func (n NaiveBoundingBox) Center() [3]float64 {
	return [3]float64{
		(n.XX - n.Xx) / 2,
		(n.YY - n.Yy) / 2,
		(n.ZZ - n.Zz) / 2,
	}
}
