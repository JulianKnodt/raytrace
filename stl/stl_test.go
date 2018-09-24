package stl

import (
	"os"
	"testing"
)

func TestBinarySTL(t *testing.T) {
	f, err := os.Open("testdata/majoras-mask.stl")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	var o []Triangle
	o, err = DecodeBinary(f)
	if err != nil {
		t.Error(err)
	}
	if len(o) == 0 {
		t.Error("Binary: Empty vertices")
	}
}

func BenchmarkBinarySTL(b *testing.B) {
	f, _ := os.Open("testdata/majoras-mask.stl")
	defer f.Close()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeBinary(f)
	}
}

func TestAsciiSTL(t *testing.T) {
	f, err := os.Open("testdata/ascii.stl")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	var o []Triangle
	o, err = DecodeAscii(f)
	if err != nil {
		t.Error(err)
	} else if len(o) == 0 {
		t.Error("Ascii: Empty Vertices")
	}
}
