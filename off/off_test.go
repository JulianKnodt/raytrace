package off

import (
	"os"
	"testing"
	/*
	  "fmt"
	  v "github.com/julianknodt/vector"
	*/)

func TestDecode(t *testing.T) {
	f, err := os.Open("./testdata/dragon.off")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = Decode(f)
	if err != nil {
		panic(err)
	}
}

/*
// Util for printing out the bounding and center of an off
func TestCenter(t *testing.T) {
	f, err := os.Open("./testdata/dragon.off")
	if err != nil {
		panic(err)
	}
	result, err := Decode(f)
	if err != nil {
		panic(err)
	}
  fmt.Println(result.Box())
  fmt.Println(*v.Average(result.Vertices))
}
*/
