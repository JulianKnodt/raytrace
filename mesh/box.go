package mesh

import (
	"math"

	"raytrace/bounding"
	v "raytrace/vector"
)

func (m Mesh) Box() bounding.Box {
	min := &v.Vec3{math.Inf(1), math.Inf(1), math.Inf(1)}
	max := &v.Vec3{math.Inf(-1), math.Inf(-1), math.Inf(-1)}
	for _, v := range m.Vertices {
		min.MinSet(v)
		max.MaxSet(v)
	}
	return bounding.Box{
		Min: *min,
		Max: *max,
	}
}