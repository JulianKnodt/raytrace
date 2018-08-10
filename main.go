package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"raytrace/camera"
	m "raytrace/material"
	"raytrace/obj"
	"raytrace/obj/mtl"
	o "raytrace/object"
	"raytrace/off"
	"raytrace/shapes"
	v "raytrace/vector"
)

var (
	width     = flag.Float64("width", 800.0, "Width to render")
	height    = flag.Float64("height", 600.0, "Height to render")
	out       = flag.String("out", "out.png", "Filepath of out file when rendering one scene")
	prof      = flag.Bool("prof", false, "Profile rendering")
	renderOpt = flag.String("render", "", "Way to render the scene")
	shift     = flag.String("shift", "0 0 0", "Shift Amount for obj files, format: \"f f f\"")
	x, y, z   float64
)

func main() {
	off := flag.String("off", "", "Off file to render")
	obj := flag.String("obj", "", "Obj file to render")
	flag.Parse()
	fmt.Sscanf(*shift, "%f %f %f", &x, &y, &z)

	if len(*off) != 0 {
		Off(*off)
		return
	} else if len(*obj) != 0 {
		Obj(*obj)
		return
	}
	run([]o.Object{
		shapes.NewSphere(v.Vec3{0, 0, -5}, 1, &m.BasicRed),
		shapes.NewSphere(v.Vec3{-1.5, 0, -5}, 1, &m.BasicBlue),
		shapes.NewTriangle(v.Vec3{0, 0, -6}, v.Vec3{2, 0, -6}, v.Vec3{1, 2, -6}, &m.BasicGreen),
		shapes.NewPlane(v.Vec3{0, -2, 0}, v.Vec3{0, 1, 0}, &m.BasicGreen),
	})
	return
}

func Off(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	model, err := off.Decode(f)
	if err != nil {
		panic(err)
	}

	model.Vertices = v.Shift(model.Vertices, x, y, z)
	run([]o.Object{model})
}

func Obj(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	model, err := obj.Decode(f, func(mtlName string) (map[string]mtl.MTL, error) {
		mtlFile, err := os.Open(filepath.Dir(filename) + mtlName)
		if err != nil {
			return nil, err
		}
		defer mtlFile.Close()

		return mtl.Decode(mtlFile)
	})
	if err != nil {
		panic(err)
	}
	model.Shift(x, y, z, 0)
	run([]o.Object{model})
}

func intersector() intersect {
	switch *renderOpt {
	case "s", "simple":
		return simpleIntersect
	default:
		return checkIntersects
	}
}

func run(o []o.Object) {
	c := camera.NewStCamera(v.Origin, camera.DefaultCameraDir, 30.0)
	l := PointLight{v.Vec3{10, 10, 10}, v.Vec3{255, 255, 255}}
	img := render(*width, *height, c, o, []Light{l}, intersector())
	file, _ := os.OpenFile("./out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	png.Encode(file, &img)
}
