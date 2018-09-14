package scene

import (
	"math"

	"raytrace/color"
	"raytrace/object"
	v "raytrace/vector"
)

func Basic(r v.Ray, s Scene) *color.Normalized {
	maxDist := math.Inf(1)
	var near object.SurfaceElement
	for _, o := range s.Objects {
		if dist, intersecting := o.Intersects(r); intersecting != nil {
			if dist < maxDist && dist > 0 {
				maxDist = dist
				near = intersecting
			}
		}
	}

	if near == nil {
		return nil
	}
	return &color.White
}
