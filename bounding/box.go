package bounding

import (
	v "github.com/julianknodt/vector"
	"math"
)

type Box struct {
	Min, Max v.Vec3
}

func NewBox(xMin, xMax, yMin, yMax, zMin, zMax float64) *Box {
	return &Box{
		Min: v.Vec3{xMin, yMin, zMin},
		Max: v.Vec3{xMax, yMax, zMax},
	}
}

func (n Box) Intersects(o Box) bool {
	return n.Max[0] > o.Min[0] && o.Max[0] > n.Min[0] &&
		n.Max[1] > o.Min[1] && o.Max[1] > n.Min[1] &&
		n.Max[2] > o.Min[2] && o.Max[2] > n.Min[2]
}

func (n Box) Contains(o Box) bool {
	return n.Max[0] > o.Max[0] && n.Min[0] < o.Min[0] &&
		n.Max[1] > o.Max[1] && n.Min[1] < o.Min[1] &&
		n.Max[2] > o.Max[2] && n.Min[2] < o.Min[2]
}

func (a Box) ContainsVec(vec v.Vec3) bool {
	return a.Min[0] < vec[0] && a.Max[0] > vec[0] &&
		a.Min[1] < vec[1] && a.Max[1] > vec[1] &&
		a.Min[2] < vec[2] && a.Max[2] > vec[2]
}

func (n Box) Center() [3]float64 {
	return [3]float64{
		(n.Max[0] + n.Min[0]) / 2,
		(n.Max[1] + n.Min[1]) / 2,
		(n.Max[2] + n.Min[2]) / 2,
	}
}

func (a Box) IntersectsRay(r v.Ray) bool {
	tMin, tMax := math.Inf(-1), math.Inf(1)

	t1 := (a.Min[0] - r.Origin[0]) / r.Direction[0]
	t2 := (a.Max[0] - r.Origin[0]) / r.Direction[0]
	tMin = math.Max(tMin, math.Min(t1, t2))
	tMax = math.Min(tMax, math.Max(t1, t2))

	t1 = (a.Min[1] - r.Origin[1]) / r.Direction[1]
	t2 = (a.Max[1] - r.Origin[1]) / r.Direction[1]
	tMin = math.Max(tMin, math.Min(t1, t2))
	tMax = math.Min(tMax, math.Max(t1, t2))

	t1 = (a.Min[2] - r.Origin[2]) / r.Direction[2]
	t2 = (a.Max[2] - r.Origin[2]) / r.Direction[2]
	tMin = math.Max(tMin, math.Min(t1, t2))
	tMax = math.Min(tMax, math.Max(t1, t2))

	return tMax > math.Max(tMin, 0)
}
