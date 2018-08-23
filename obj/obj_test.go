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
	defer f.Close()

	if _, err = Decode(f); err != nil {
		t.Error(err)
	}
}

func TestCube(t *testing.T) {
	f, err := os.Open("./testdata/cube.obj")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	if _, err = Decode(f); err != nil {
		t.Error(err)
	}
}

func TestTeapotComplex(t *testing.T) {
	f, err := os.Open("./testdata/teapot/teapot.obj")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	// there's no need to panic here since it should propogate upwards.
	if _, err = Decode(f); err != nil {
		t.Error(err)
	}
}
