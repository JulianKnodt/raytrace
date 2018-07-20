package main

import (
	obj "github.com/julianknodt/raytrace/object"
	v "github.com/julianknodt/raytrace/vector"
	"image"
	"image/color"
	"math"
	"runtime"
)

const epsilon = 1e-6

func checkIntersects(from, dir v.Vec3, objects []obj.Object, lights []Light) color.Color {
	maxDist := math.Inf(1)
	var near obj.Shape
	for _, o := range objects {
		if dist, intersecting := o.Intersects(from, dir); intersecting != nil {
			if dist < maxDist && dist > 0 {
				maxDist = dist
				near = intersecting
			}
		}
	}

	if near == nil {
		return color.Black
	}

	if maxDist < 0 {
		panic("something behind the camera is supposed to be visible")
	}

	inter := v.Add(from, v.SMul(maxDist, dir))
	normalInter, invAble := near.Normal(inter)
	v.UnitSet(&normalInter)
	v.AddSet(&inter, v.SMul(epsilon, normalInter))

	var color v.Vec3
	for _, l := range lights {
		lightDir, canIllum := l.LightTo(inter) // intersection -> light
		align := v.Dot(normalInter, lightDir)
		if !canIllum || align <= 0 {
			if align < 0 && invAble {
				align = -align
				v.AddSet(&inter, v.SMul(-2*epsilon, normalInter))
			} else {
				continue
			}
		}
		for _, o := range objects {
			if _, intersecting := o.Intersects(inter, lightDir); intersecting != nil {
				canIllum = false
				break
			}
		}
		if canIllum {
			v.AddSet(&color, v.SMul(align, near.Color()))
		}
	}
	return v.ToRGBA(color)
}

type coord struct {
	x float64
	y float64
}

type fieldColor struct {
	x     int
	y     int
	color color.Color
}

func render(width, height float64, cam Camera, o []obj.Object, l []Light) image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	var invWidth float64 = 1.0 / width
	var invHeight float64 = 1.0 / height
	aspectRatio := width * invHeight
	angle := math.Tan(math.Pi * 0.5 * cam.FOV() / 180)
	out := make(chan fieldColor, int(height*width))
	work := make(chan coord)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for c := range work {
				xDir := (2*((c.x+0.5)*invWidth) - 1) * angle * aspectRatio
				yDir := (1 - 2*((c.y+0.5)*invHeight)) * angle
				direction := v.Unit(v.Sub(v.Vec3{xDir, yDir, -1}, cam.Location()))
				out <- fieldColor{int(c.x), int(c.y), checkIntersects(v.Origin, direction, o, l)}
			}
		}()
	}
	for y := 0.0; y < height; y++ {
		for x := 0.0; x < width; x++ {
			work <- coord{x, y}
		}
	}
	close(work)

	for i := 0; i < cap(out); i++ {
		o := <-out
		img.Set(o.x, o.y, o.color)
	}
	close(out)

	return *img
}
