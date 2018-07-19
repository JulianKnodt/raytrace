package vector

func Normal(v []Vec3) Vec3 {
	if !Coplanar(v) {
		panic("not coplanar")
	}
	return Cross(Sub(v[0], v[1]), Sub(v[0], v[2]))
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
	// todo
	return false
}
