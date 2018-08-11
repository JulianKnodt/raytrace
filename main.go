package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"raytrace/camera"
	"raytrace/light"
	"raytrace/scene"
	v "raytrace/vector"
)

// 🎌
var (
	width     = flag.Float64("width", 800.0, "Width to render")
	height    = flag.Float64("height", 600.0, "Height to render")
	out       = flag.String("out", "out.png", "Filepath of out file when rendering one scene")
	prof      = flag.Bool("prof", false, "Profile rendering")
	renderOpt = flag.String("render", "", "Way to render the scene")
	shift     = flag.String("shift", "0 0 0", "Shift Amount for obj files, format: \"f f f\"")
	skyFile   = flag.String("sky", "", "Image to be used as the sky")
)

// global variables for this scope
var (
	x, y, z float64
)

func main() {
	off := flag.String("off", "", "Off file to render")
	obj := flag.String("obj", "", "Obj file to render")
	flag.Parse()
	fmt.Sscanf(*shift, "%f %f %f", &x, &y, &z)

	scene := scene.Scene{
		Height:               *height,
		Width:                *width,
		IntersectionFunction: intersector(),
		Camera:               camera.NewStCamera(v.Origin, camera.DefaultCameraDir, 30.0),
		Lights: []light.Light{
			light.PointLight{RadiantColor: v.Vec3{10, 10, 10}, Center: v.Vec3{255, 255, 255}},
		},
	}

	scene.AddObj(*obj, v.Vec3{x, y, z})
	scene.AddOff(*off, v.Vec3{x, y, z})
	scene.AddSky(*skyFile)
	img := scene.Render()
	file, _ := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	png.Encode(file, &img)
}

func intersector() scene.Intersector {
	switch *renderOpt {
	case "s", "simple":
		return scene.Basic
	default:
		return scene.Direct
	}
}
