package obj

import (
	"os"
	"testing"
)

func TestObj(t *testing.T) {
	f, err := os.Open("./teapot.obj")
	if err != nil {
		t.Error(err)
	}

	_, err = Decode(f)
	if err != nil {
		t.Error(err)
	}
}
