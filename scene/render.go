package scene

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"raytrace/light"
	obj "raytrace/object"
	v "raytrace/vector"
	"runtime"
	"sync/atomic"
	"time"
)

const epsilon = 1e-6

type intersect func(v.Vec3, v.Vec3, []obj.Object, []light.Light) color.Color

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
func (s Scene) Render() image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, int(s.Width), int(s.Height)))
	var invWidth float64 = 1.0 / s.Width
	var invHeight float64 = 1.0 / s.Height
	aspectRatio := s.Width * invHeight
	angle := math.Tan(math.Pi * 0.5 * s.Camera.FOV() / 180)
	out := make(chan fieldColor, int(s.Height*s.Width))
	work := make(chan coord, int(s.Height*s.Width))
	count := int64(0)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for c := range work {
				xDir := (2*((c.x+0.5)*invWidth) - 1) * angle * aspectRatio
				yDir := (1 - 2*((c.y+0.5)*invHeight)) * angle
				direction := v.Unit(v.Sub(v.Vec3{xDir, yDir, -1}, s.Camera.Location()))
				out <- fieldColor{
					int(c.x),
					int(c.y),
					s.IntersectionFunction(v.Origin, direction, s),
				}
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	for y := 0.0; y < s.Height; y++ {
		for x := 0.0; x < s.Width; x++ {
			work <- coord{x, y}
		}
	}
	close(work)

	timer := time.NewTimer(5 * time.Second)
	start := time.Now()
	for count < int64(s.Height*s.Width) {
		select {
		case o := <-out:
			img.Set(o.x, o.y, o.color)
		case <-timer.C:
			fmt.Printf(
				"Time elapsed %s | Percent Done %.3f%% | Pixels Complete %d/%d\n",
				time.Since(start).Round(time.Second),
				float64(count)/(s.Height*s.Width)*100,
				count,
				int(s.Height*s.Width),
			)
			timer.Reset(5 * time.Second)
		}
	}
	close(out)
	for o := range out {
		img.Set(o.x, o.y, o.color)
	}

	return *img
}