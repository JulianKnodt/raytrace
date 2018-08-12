package main

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	num := 10
	radius := 5.0
	c := Circle(radius, num)
	if len(c) != num {
		t.Fail()
	}
	for _, point := range c {
		if math.Round(math.Hypot(point[0], point[1])) != radius {
			t.Fail()
		}
	}
}
