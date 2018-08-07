package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	obj "raytrace/object"
	v "raytrace/vector"
	"runtime"
	"sync/atomic"
	"time"
)

const epsilon = 1e-6

type intersect func(v.Vec3, v.Vec3, []obj.Object, []Light) color.Color

func simpleIntersect(origin, dir v.Vec3, objects []obj.Object, lights []Light) color.Color {
	maxDist := math.Inf(1)
	var near obj.SurfaceElement
	for _, o := range objects {
		if dist, intersecting := o.Intersects(origin, dir); intersecting != nil {
			if dist < maxDist && dist > 0 {
				maxDist = dist
				near = intersecting
			}
		}
	}

	if near == nil {
		return color.Black
	}
	return color.White
}

func checkIntersects(from, dir v.Vec3, objects []obj.Object, lights []Light) color.Color {
	maxDist := math.Inf(1)
	var near obj.SurfaceElement
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
			v.AddSet(&color, v.SMul(align, near.Mat().Emitted()))
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

// Returns an image of given height and width
// with the given objects using 'inter' intersection
// algorithm choice
func render(
	width, height float64,
	cam Camera,
	o []obj.Object,
	l []Light,
	inter intersect,
) image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	var invWidth float64 = 1.0 / width
	var invHeight float64 = 1.0 / height
	aspectRatio := width * invHeight
	angle := math.Tan(math.Pi * 0.5 * cam.FOV() / 180)
	out := make(chan fieldColor, int(height*width))
	work := make(chan coord, int(height*width))
	count := int64(0)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for c := range work {
				xDir := (2*((c.x+0.5)*invWidth) - 1) * angle * aspectRatio
				yDir := (1 - 2*((c.y+0.5)*invHeight)) * angle
				direction := v.Unit(v.Sub(v.Vec3{xDir, yDir, -1}, cam.Location()))
				out <- fieldColor{int(c.x), int(c.y), inter(v.Origin, direction, o, l)}
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	for y := 0.0; y < height; y++ {
		for x := 0.0; x < width; x++ {
			work <- coord{x, y}
		}
	}
	close(work)

	timer := time.NewTimer(5 * time.Second)
	start := time.Now()
	for count < int64(height*width) {
		select {
		case o := <-out:
			img.Set(o.x, o.y, o.color)
		case <-timer.C:
			fmt.Printf("Time elapsed %s | Percent Done %.3f%% | Pixels Complete %d/%d\n", time.Since(start).Round(time.Second), float64(count)/(height*width)*100, count, int(height*width))
			timer.Reset(5 * time.Second)
		}
	}
	close(out)
	for o := range out {
		img.Set(o.x, o.y, o.color)
	}

	return *img
}
