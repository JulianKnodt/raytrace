package obj

import (
	"os"
	"testing"
)

func TestTeapotBasic(t *testing.T) {
	f, err := os.Open("./testdata/teapot.obj")
	if err != nil {
		t.Error(err)
	}

	_, err = Decode(f)
	if err != nil {
		t.Error(err)
	}
}

func TestCube(t *testing.T) {
	f, err := os.Open("./testdata/cube.obj")
	if err != nil {
		t.Error(err)
	}

	_, err = Decode(f)
	if err != nil {
		t.Error(err)
	}
}
