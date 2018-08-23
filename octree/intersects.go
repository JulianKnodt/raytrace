package octree

import (
	"math"
	"raytrace/object"
	v "raytrace/vector"
)

var posInf = math.Inf(1)

func (o Octree) Intersects(r v.Ray) (float64, object.SurfaceElement) {
	if o.Region.IntersectsRay(r) {
		min := posInf
		var res object.SurfaceElement
		for _, child := range o.Children {
			if child == nil {
				break
			} else if t, surfel := child.Intersects(r); t < min {
				min = t
				res = surfel
			}
		}

		for _, val := range o.processedValues {
			if t, surfel := val.Intersects(r); t < min && t > 0 {
				min = t
				res = surfel
			}
		}

		return min, res
	} else {
		return posInf, nil
	}
}
