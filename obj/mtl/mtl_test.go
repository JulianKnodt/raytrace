package mtl

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	files, err := ioutil.ReadDir("./testdata/")
	if err != nil {
		t.Error(err)
	}
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".mtl") {
			continue
		}
		file, err := os.Open(fmt.Sprintf("./testdata/%s", f.Name()))
		if err != nil {
			t.Error(err)
		}
		mtls, err := Decode(file)
		if err != nil {
			t.Error(err)
		}

		for _, v := range mtls {
			v.Material()
		}

	}
}
