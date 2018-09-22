package main

import (
	"flag"
	"image/png"
	"math/rand"
	"os"
	"time"

	v "github.com/julianknodt/vector"
	"raytrace/camera"
	"raytrace/object"
	"raytrace/scene"
	"raytrace/shapes"
)

var rs = rand.New(rand.NewSource(time.Now().UnixNano()))

const numItems = 1000

var (
	out = flag.String("o", "out.png", "Where to put output image")
)

func main() {
	objects := make([]object.Object, 0, numItems)
	for i := 0; i < numItems; i++ {
		objects = append(objects, shapes.NewSphere(v.RandomVector(), rs.Float64(), nil))
	}
	scene := scene.Scene{
		Height:               1000.0,
		Width:                1200.0,
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
