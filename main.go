package main

import (
	"flag"
	"github.com/julianknodt/raytrace/off"
	"os"
)

func _main() {
	off := flag.String("off", "", "Off file to render")

	if len(*off) != 0 {
		Off(*off)
	}
}

func Off(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	off.Decode(f)
}
