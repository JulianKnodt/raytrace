package main

import (
	"flag"
	"image/png"
	"os"

	v "github.com/julianknodt/vector"
	"raytrace/camera"
	"raytrace/object"
	"raytrace/scene"
	"raytrace/shapes"
)

var (
	out = flag.String("o", "out.png", "Where to put output image")
)

func main() {
	objects := []object.Object{
		shapes.NewSphere(v.Vec3{0, 0, -5}, 4, nil),
	}
	scene := scene.Scene{
		Height:               1000.0,
		Width:                2000.0,
		IntersectionFunction: scene.Direct,
		Camera:               camera.DefaultCamera(),
		Objects:              objects,
	}

	img := scene.Render()
	f, err := os.Create(*out)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
