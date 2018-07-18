package main

import (
	v "github.com/julianknodt/raytrace/vector"
	"image/png"
	"os"
)

func main() {
	width := 800.0
	height := 600.0
	s := NewSphere(v.Vec3{0, 0, -5}, 1, v.Vec3{255, 0, 0})
	s2 := NewSphere(v.Vec3{-1.5, 0, -5}, 1, v.Vec3{0, 0, 255})

	t := NewTriangle(v.Vec3{0, 0, -6}, v.Vec3{2, 0, -6}, v.Vec3{1, 2, -6}, v.Vec3{122, 122, 122})

	p := NewPlane(v.Vec3{0, -2, 0}, v.Vec3{0, 1, 0}, v.Vec3{0, 255, 0})
	c := NewStCamera(v.Origin, DefaultCameraDir, 30.0)
	l := PointLight{v.Vec3{10, 10, 10}, v.Vec3{255, 255, 255}}
	img := render(width, height, c, []Object{*t, *s, *s2, *p}, []Light{l})
	file, _ := os.OpenFile("./out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	png.Encode(file, &img)
}
