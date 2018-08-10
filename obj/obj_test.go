package obj

import (
	"os"
	"raytrace/obj/mtl"
	"testing"
)

func TestTeapotBasic(t *testing.T) {
	f, err := os.Open("./testdata/teapot.obj")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	_, err = Decode(f, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestCube(t *testing.T) {
	f, err := os.Open("./testdata/cube.obj")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	_, err = Decode(f, nil)
	if err != nil {
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
	_, err = Decode(f, func(name string) (map[string]mtl.MTL, error) {
		mtlFile, err := os.Open("./testdata/teapot/" + name)
		if err != nil {
			return nil, err
		}
		defer mtlFile.Close()
		mtls, err := mtl.Decode(mtlFile)
		if err != nil {
			return nil, err
		}
		return mtls, nil
	})

	if err != nil {
		t.Error(err)
	}
}
