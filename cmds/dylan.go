package main

import (
	"image/png"
	"os"

	"raytrace/camera"
	"raytrace/color"
	"raytrace/light"
	"raytrace/material"
	"raytrace/scene"
	"raytrace/shapes"
	"raytrace/utils"
	v "raytrace/vector"
)

func main() {
	scene := scene.Scene{
		Height:               1200.0,
		Width:                1600.0,
		IntersectionFunction: scene.Direct,
		Camera:               camera.DefaultCamera(),
		Lights: []light.Light{
			light.PointLight{RadiantColor: v.Vec3{10, 10, 10}, Center: v.Vec3{255, 255, 255}},
		},
	}
	img, err := utils.LoadImage("../samples/dylan.jpg")
	if err != nil {
		panic(err)
	}
	scene.Objects = append(scene.Objects, shapes.NewSphere(v.Vec3{0, 0, -50}, 10,
		&material.Material{
			Ambient:        color.Blank,
			AmbientTexture: img,
		}))

	out := scene.Render()

	f, err := os.Create("./dylan.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = png.Encode(f, out); err != nil {
		panic(err)
	}
}
