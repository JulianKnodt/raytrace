package camera

import (
	"fmt"
	"testing"
)

func TestTo(t *testing.T) {
	fmt.Println(DefaultCamera().To(0, 0))
}
