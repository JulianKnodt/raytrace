package main

import (
	"image/png"
	"os"
	m "raytrace/material"
	obj "raytrace/object"
	"raytrace/shapes"
	v "raytrace/vector"
)

func _main() {
	width := 800.0
	height := 600.0
	s := shapes.NewSphere(v.Vec3{0, 0, -5}, 1, &m.BasicRed)
	s2 := shapes.NewSphere(v.Vec3{-1.5, 0, -5}, 1, &m.BasicBlue)

	t := shapes.NewTriangle(v.Vec3{0, 0, -6}, v.Vec3{2, 0, -6}, v.Vec3{1, 2, -6}, &m.BasicGreen)

	p := shapes.NewPlane(v.Vec3{0, -2, 0}, v.Vec3{0, 1, 0}, &m.BasicGreen)
	c := NewStCamera(v.Origin, DefaultCameraDir, 30.0)
	l := PointLight{v.Vec3{10, 10, 10}, v.Vec3{255, 255, 255}}
	img := render(width, height, c, []obj.Object{*t, *s, *s2, *p}, []Light{l})
	file, _ := os.OpenFile("./out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	png.Encode(file, &img)
}
