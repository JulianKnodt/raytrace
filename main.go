package main

import (
	"flag"
	"github.com/julianknodt/raytrace/off"
	"os"
)

func _main() {
	off := flag.String("off", "", "Off file to render")
  width := flag.Float64("width", 800.0, "Width to render")
  height := flag.Float64("height", 600.0, "Height to render")
  out := flag.String("out", "out.png", "Filepath of out file when rendering one scene")

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
