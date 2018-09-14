package obj

import (
	"encoding/json"
	"os"
	"testing"
)

func TestIndexedTriangleList(t *testing.T) {
	f, err := os.Open("./testdata/teapot.obj")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	obj, err := Decode(f)
	if err != nil {
		t.Error(err)
	}

	_, err = json.Marshal(obj.IndexedTriangleList())
	if err != nil {
		t.Error(err)
	}
}
