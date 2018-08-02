package main

import (
	"flag"
	"os"
	"raytrace/off"
	//	"raytrace/obj" // oops deleted this from git by accident
)

var (
	width  = flag.Float64("width", 800.0, "Width to render")
	height = flag.Float64("height", 600.0, "Height to render")
	out    = flag.String("out", "out.png", "Filepath of out file when rendering one scene")
)

func _main() {
	off := flag.String("off", "", "Off file to render")
	obj := flag.String("obj", "", "Obj file to render")

	flag.Parse()

	if len(*off) != 0 {
		Off(*off)
	} else if len(*obj) != 0 {
		Obj(*obj)
	}
}

func Off(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	off.Decode(f)
}

func Obj(filename string) {
	/*
	   f, err := os.Open(filename)
	   if err != nil {
	     panic(err)
	   }

	   obj.Decode(f)
	*/
}
