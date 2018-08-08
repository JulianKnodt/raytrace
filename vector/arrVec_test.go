package vector

import (
	"testing"
)

func TestShift(t *testing.T) {
	a := []Vec3{Vec3{0, 0, 0}}
	a = Shift(a, 1, 1, 1)
	if a[0][0] != 1 {
		t.Error("Did not shift correctly")
	}
}
