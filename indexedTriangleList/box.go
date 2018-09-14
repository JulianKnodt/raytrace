package indexedTriangleList

import (
	"math"

	v "github.com/julianknodt/vector"
	"raytrace/bounding"
)

func (i IndexedTriangleList) Box() bounding.Box {
	min := &v.Vec3{math.Inf(1), math.Inf(1), math.Inf(1)}
	max := &v.Vec3{math.Inf(-1), math.Inf(-1), math.Inf(-1)}
	for _, v := range i.Vertices {
		min.MinSet(v)
		max.MaxSet(v)
	}
	return bounding.Box{
		Min: *min,
		Max: *max,
	}
}
