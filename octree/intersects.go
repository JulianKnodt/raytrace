package octree

import (
	"math"
	"raytrace/object"
	v "raytrace/vector"
)

var posInf = math.Inf(1)

func (o Octree) Intersects(origin, dir v.Vec3) (float64, object.SurfaceElement) {
	if o.Region.IntersectsRay(origin, dir) {
		min := posInf
		var res object.SurfaceElement
		for _, child := range o.Children {
			if child == nil {
				break
			} else if t, surfel := child.Intersects(origin, dir); t < min {
				min = t
				res = surfel
			}
		}

		for _, val := range o.processedValues {
			if t, surfel := val.Intersects(origin, dir); t < min && t > 0 {
				min = t
				res = surfel
			}
		}

		return min, res
	} else {
		return posInf, nil
	}
}
