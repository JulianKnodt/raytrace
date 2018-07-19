package off

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("./dragon.off")
	if err != nil {
		panic(err)
	}
	_, err = Decode(f)
	if err != nil {
		panic(err)
	}
}
