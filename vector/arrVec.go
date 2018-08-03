package vector

import (
	"math"
)

func normalNoCheck(v []Vec3) (Vec3, bool) {
	return Cross(Sub(v[0], v[1]), Sub(v[0], v[2])), true
}

func Normal(v []Vec3) (Vec3, bool) {
	if !Coplanar(v) {
		return Vec3{}, false
	}
	return Cross(Sub(v[0], v[1]), Sub(v[0], v[2])), true
}

func Coplanar(v []Vec3) bool {
	if len(v) <= 3 {
		return true
	}

	test := Cross(Sub(v[0], v[1]), Sub(v[0], v[2]))

	for _, vec := range v[2:] {
		if Dot(test, Sub(v[0], vec)) != 0 {
			return false
		}
	}
	return true
}

func Colinear(v []Vec3) bool {
	if len(v) <= 2 {
		return true
	}

	test := Sub(v[0], v[1])
	for _, vec := range v[2:] {
		if !RelEqual(test, Sub(v[0], vec)) {
			return false
		}
	}
	return true
}

const epsilon = 0.000001

func IntersectsTriangle(a, b, c, origin, dir Vec3) (float64, bool) {
	edge1 := Sub(b, a)
	edge2 := Sub(c, a)
	h := Cross(dir, edge2)
	area := Dot(edge1, h)
	if area > -epsilon && area < epsilon {
		return -1, false // this is collinear
	}

	invArea := 1 / area
	s := Sub(origin, a)
	u := invArea * Dot(s, h)
	if u < 0 || u > 1 {
		return -1, false
	}

	q := Cross(s, edge1)
	v := invArea * Dot(dir, q)
	if v < 0 || (u+v) > 1 {
		return -1, false
	}

	par := invArea * Dot(edge2, q)
	if par > epsilon {
		return par, true
	}
	return par, false
}

func Intersects(v []Vec3, origin, dir Vec3) (float64, bool) {
	for i, vec := range v[:len(v)-2] {
		if t, intersects := IntersectsTriangle(vec, v[i+1], v[i+2], origin, dir); intersects {
			return t, true
		}
	}
	return -1, false
}

func Interpolate(points [3]Vec3, barycentric Vec3) Vec3 {
	result := Vec3{}
	for i, point := range points {
		AddSet(&result, SMul(barycentric[i], point))
	}
	return result
}

func IntersectsTriangle2(a, b, c, origin, dir Vec3) (float64, bool, Vec3) {
	edge1 := Sub(b, a)
	edge2 := Sub(c, a)

	n := Unit(Cross(edge1, edge2)) // normal to triangle

	q := Cross(dir, edge2)
	p := Dot(edge1, q)

	if Dot(n, dir) >= 0 || math.Abs(p) <= epsilon {
		return -1, false, Vec3{}
	}

	s := Op(Sub(origin, a), func(px float64) float64 {
		return px / p
	})
	r := Cross(s, edge1)

	br := Vec3{}
	br[0] = Dot(s, q)
	br[1] = Dot(r, dir)
	br[2] = 1 - br[0] - br[1]

	if br[0] < 0 || br[1] < 0 || br[2] < 0 {
		return -1, false, br
	}

	t := Dot(edge2, r)
	return t, t >= 0, br
}
