package utils

import (
	"math"
)

func Maxmin(a, b, c float64) (max, min float64) {
	switch {
	case a >= b && a >= c:
		return a, math.Min(b, c)
	case b >= a && b >= c:
		return b, math.Min(a, c)
	case c >= a && c >= b:
		return c, math.Min(b, a)
	default:
		panic("Somehow there is not a min/max amidst three floats")
	}
}
