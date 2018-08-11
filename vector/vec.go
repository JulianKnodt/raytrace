package vector

import (
	"image/color"
	"math"
)

type Vec3 [3]float64

var Origin Vec3 = Vec3{0, 0, 0}

func NewOrigin() *Vec3 {
	return &Vec3{0, 0, 0}
}

func NewVec(a float64) *Vec3 {
	return &Vec3{a, a, a}
}

func (v Vec3) X() float64 {
	return v[0]
}

func (v Vec3) Y() float64 {
	return v[1]
}

func (v Vec3) Z() float64 {
	return v[2]
}

func Equal(a, b Vec3) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}

func RelEqual(a, b Vec3) bool {
	firstRel := a[0] / b[0]
	return firstRel == a[1]/b[1] && firstRel == a[2]/b[2]
}

func Sub(a, b Vec3) (r Vec3) {
	r[0] = a[0] - b[0]
	r[1] = a[1] - b[1]
	r[2] = a[2] - b[2]
	return
}

func InvSet(a *Vec3) *Vec3 {
	a[0] = -a[0]
	a[1] = -a[1]
	a[2] = -a[2]
	return a
}

func SqrMagn(a Vec3) float64 {
	return a[0]*a[0] + a[1]*a[1] + a[2]*a[2]
}

func Magn(a Vec3) float64 {
	return math.Sqrt(SqrMagn(a))
}

func Inv(a Vec3) Vec3 {
	return Vec3{-a[0], -a[1], -a[2]}
}

func Dot(a, b Vec3) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func sqr(a float64) float64 {
	return a * a
}

// Returns radians between a and b
func Theta(a, b Vec3) float64 {
	return math.Acos(Dot(a, b) / (Magn(a) * Magn(b)))
}

func PMul(a, b Vec3) (r Vec3) {
	r[0] = a[0] * b[0]
	r[1] = a[1] * b[1]
	r[2] = a[2] * b[2]
	return
}

func DistSquared(a, b Vec3) float64 {
	return sqr(a[0]-b[0]) + sqr(a[1]-b[1]) + sqr(a[2]-b[2])
}

func ToRGBA(a Vec3) color.RGBA {
	return color.RGBA{uint8(a[0]), uint8(a[1]), uint8(a[2]), 255}
}

func Cross(a, b Vec3) Vec3 {
	return Vec3{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

func Op(in Vec3, op func(float64) float64) (out Vec3) {
	for i, v := range in {
		out[i] = op(v)
	}
	return
}
