package main

import (
	"flag"
	"image/png"
	"math/rand"
	"os"
	"time"

	v "github.com/julianknodt/vector"
	"raytrace/camera"
	"raytrace/color"
	"raytrace/material"
	"raytrace/object"
	"raytrace/scene"
	"raytrace/shapes"
)

var rs = rand.New(rand.NewSource(time.Now().UnixNano()))

const numItems = 100

var (
	out = flag.String("o", "out.png", "Where to put output image")
)

func Sub(a float64) float64 {
	return a - 0.5
}

func main() {
	flag.Parse()
	objects := make([]object.Object, 0, numItems+1)
	for i := 0; i < numItems; i++ {
		objects = append(
			objects,
			shapes.NewSphere(
				*v.RandomVector().
					OpSet(Sub).
					SMulSet(10), rs.Float64(), &material.Material{
					Reflect:    true,
					Ambient:    color.Normalized{RGB: v.Vec3{rs.Float64(), rs.Float64(), rs.Float64()}, A: 1},
					RenderType: material.Lambertian{Albedo: 1},
				}),
		)
	}
	objects = append(objects, shapes.NewSphere(v.Vec3{0, 0, 0}, 0.5, nil))
	lights := []object.LightSource{
		shapes.NewSphere(v.Vec3{-12, 0, -12}, 1, material.WhiteLightMaterial()),
	}

	c := camera.DefaultCamera()
	c.Transform.Origin[2] -= -6
	scene := scene.Scene{
		Height:               1000.0,
		Width:                1000.0,
		IntersectionFunction: scene.Direct,
		Camera:               c,
		Objects:              objects,
		Lights:               lights,
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
