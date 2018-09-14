package quaternion

import (
	"math"

	v "raytrace/vector"
)

// Represents
type Quaternion struct {
	Scalar float64
	Vec    v.Vec3
}

func (q Quaternion) Mul(o Quaternion) *Quaternion {
	return &Quaternion{
		q.Scalar*o.Scalar - q.Vec.Dot(o.Vec),
		*q.Vec.SMul(o.Scalar).
			AddSet(*o.Vec.SMul(q.Scalar)).
			AddSet(*o.Vec.Cross(q.Vec)),
	}
}

func (q Quaternion) Conj() *Quaternion {
	return &Quaternion{
		q.Scalar,
		*q.Vec.Inv(),
	}
}

func sqr(a float64) float64 {
	return a * a
}

func (q Quaternion) SqrLength() float64 {
	return sqr(q.Scalar) + q.Vec.SqrMagn()
}

func (q Quaternion) Length() float64 {
	return math.Sqrt(q.SqrLength())
}

func (q Quaternion) Inverse() *Quaternion {
	s := q.SqrLength()
	out := q.Conj()
	out.Scalar /= s
	(&out.Vec).SMulSet(1 / s)
	return out
}
