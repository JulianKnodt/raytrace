package obj

import (
	"encoding/json"
	"fmt"
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

	b, err := json.Marshal(obj.IndexedTriangleList())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
}
