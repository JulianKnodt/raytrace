package scene

import (
	"fmt"
	v "github.com/julianknodt/vector"
	"image"
	"image/color"
	obj "raytrace/object"
	"runtime"
	"sync/atomic"
	"time"
)

const epsilon = 1e-6

type intersect func(v.Vec3, v.Vec3, []obj.Object, []obj.Object) color.Color

type coord struct {
	x int
	y int
}

type fieldColor struct {
	x     int
	y     int
	color color.Color
}

// Returns an image of given height and width
// with the given objects using 'inter' intersection
// algorithm choice
func (s Scene) Render() *image.RGBA {
	switch {
	case s.IntersectionFunction == nil:
		panic("Nil intersection function")
	}
	img := image.NewRGBA(image.Rect(0, 0, int(s.Width), int(s.Height)))
	out := make(chan fieldColor, int(s.Height*s.Width))
	work := make(chan coord, int(s.Height*s.Width))
	count := int64(0)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for c := range work {
				out <- fieldColor{
					c.x, c.y,
					s.IntersectionFunction(
						s.Camera.RayTo(float64(c.x)/s.Width, float64(c.y)/s.Height),
						s,
					).ToImageColor(),
				}
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	for y := 0.0; y < s.Height; y++ {
		for x := 0.0; x < s.Width; x++ {
			work <- coord{int(x), int(y)}
		}
	}
	close(work)

	timer := time.NewTimer(5 * time.Second)
	start := time.Now()
	for count < int64(s.Height*s.Width) {
		select {
		case o := <-out:
			if o.color == nil {
				o.color = s.Sky.At(o.x, o.y)
			}
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
		if o.color == nil {
			o.color = s.Sky.At(o.x, o.y)
		}
		img.Set(o.x, o.y, o.color)
	}

	return img
}
