package material

import (
	"fmt"
	"raytrace/utils"
	"testing"
)

func TestSample(t *testing.T) {
	img, err := utils.LoadImage("./testdata/dog.jpg")
	if err != nil {
		t.Error(err)
	}

	it := ImageTexture{img}

	fmt.Println(it.Sample(1, 1))
}
