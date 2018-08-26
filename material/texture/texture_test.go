package material

import (
	"raytrace/utils"
	"testing"
)

func TestSample(t *testing.T) {
	img, err := utils.LoadImage("./testdata/dog.jpg")
	if err != nil {
		t.Error(err)
	}

	Sample(img, 1, 1)
}
