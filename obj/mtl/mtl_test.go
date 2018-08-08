package mtl

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	files, err := ioutil.ReadDir("./testdata/")
	if err != nil {
		t.Error(err)
	}
	for _, f := range files {
		file, err := os.Open(fmt.Sprintf("./testdata/%s", f.Name()))
		if err != nil {
			t.Error(err)
		}
		out, err := Decode(file)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(out)

	}
}
