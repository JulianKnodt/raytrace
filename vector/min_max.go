package vector

import (
	"math"
)

func (v Vec3) Max(o Vec3) Vec3 {
	return Vec3{
		math.Max(v[0], o[0]),
		math.Max(v[1], o[1]),
		math.Max(v[2], o[2]),
	}
}

func Max(v, o Vec3) Vec3 {
	return Vec3{
		math.Max(v[0], o[0]),
		math.Max(v[1], o[1]),
		math.Max(v[2], o[2]),
	}
}

func (v Vec3) Min(o Vec3) Vec3 {
	return Vec3{
		math.Min(v[0], o[0]),
		math.Min(v[1], o[1]),
		math.Min(v[2], o[2]),
	}
}

func Min(v, o Vec3) Vec3 {
	return Vec3{
		math.Min(v[0], o[0]),
		math.Min(v[1], o[1]),
		math.Min(v[2], o[2]),
	}
}
