package bounding

import (
	"math"

	v "github.com/julianknodt/vector"
)

// A box which is aligned with the axis
// intended to be used to test whether or not something intersects it
type AxisAlignedBoundingBox struct {
	// Min x
	Xx float64
	// Max X
	XX float64
	// Min y
	Yy float64
	// Max y
	YY float64
	// Min z
	Zz float64
	// Max z
	ZZ float64
}

func NewOriginAABB(size float64) *AxisAlignedBoundingBox {
	size = math.Abs(size)
	return &AxisAlignedBoundingBox{
		-size, size, -size, size, -size, size,
	}
}

func (n AxisAlignedBoundingBox) Intersects(o AxisAlignedBoundingBox) bool {
	return n.XX > o.Xx && o.XX > n.Xx &&
		n.YY > o.Yy && o.YY > n.Yy &&
		n.ZZ > o.Zz && o.ZZ > n.Zz
}

func (n AxisAlignedBoundingBox) Contains(o AxisAlignedBoundingBox) bool {
	return n.XX > o.XX && n.Xx < o.Xx &&
		n.YY > o.YY && n.Yy < o.Yy &&
		n.ZZ > o.ZZ && n.Zz < o.Zz
}

func (a AxisAlignedBoundingBox) ContainsVec(vec v.Vec3) bool {
	return a.Xx < vec[0] && a.XX > vec[0] &&
		a.Yy < vec[1] && a.YY > vec[1] &&
		a.Zz < vec[2] && a.ZZ > vec[2]
}

func (n AxisAlignedBoundingBox) Center() [3]float64 {
	return [3]float64{
		(n.XX + n.Xx) / 2,
		(n.YY + n.Yy) / 2,
		(n.ZZ + n.Zz) / 2,
	}
}

func (a AxisAlignedBoundingBox) IntersectsRay(r v.Ray) bool {
	tMin, tMax := math.Inf(-1), math.Inf(1)

	t1 := (a.Xx - r.Origin[0]) / r.Direction[0]
	t2 := (a.XX - r.Origin[0]) / r.Direction[0]
	tMin = math.Max(tMin, math.Min(t1, t2))
	tMax = math.Min(tMax, math.Max(t1, t2))

	t1 = (a.Yy - r.Origin[1]) / r.Direction[1]
	t2 = (a.YY - r.Origin[1]) / r.Direction[1]
	tMin = math.Max(tMin, math.Min(t1, t2))
	tMax = math.Min(tMax, math.Max(t1, t2))

	t1 = (a.Zz - r.Origin[2]) / r.Direction[2]
	t2 = (a.ZZ - r.Origin[2]) / r.Direction[2]
	tMin = math.Max(tMin, math.Min(t1, t2))
	tMax = math.Min(tMax, math.Max(t1, t2))

	return tMax > math.Max(tMin, 0)
}
