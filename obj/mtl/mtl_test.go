package mtl

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("./testdata/flat_green.mtl")
	if err != nil {
		t.Error(err)
	}

	_, err = Decode(f)
	if err != nil {
		t.Error(err)
	}
}
