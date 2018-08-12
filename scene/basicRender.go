package scene

import (
	"image/color"
	"math"
	"raytrace/object"
	v "raytrace/vector"
)

func Basic(origin, dir v.Vec3, s Scene) color.Color {
	maxDist := math.Inf(1)
	var near object.SurfaceElement
	for _, o := range s.Objects {
		if dist, intersecting := o.Intersects(origin, dir); intersecting != nil {
			if dist < maxDist && dist > 0 {
				maxDist = dist
				near = intersecting
			}
		}
	}

	if near == nil {
		return nil
	}
	return color.White
}
