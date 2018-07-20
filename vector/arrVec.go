package vector

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
