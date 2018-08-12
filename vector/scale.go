package vector

func Scale(dst Vec3, x, y, z float64) (r Vec3) {
	r[0] = dst[0] * x
	r[1] = dst[1] * y
	r[2] = dst[2] * z
	return
}

func ScaleSet(dst *Vec3, x, y, z float64) {
	dst[0] = dst[0] * x
	dst[1] = dst[1] * y
	dst[2] = dst[2] * z
}

func (v Vec3) Scale(x, y, z float64) Vec3 {
	return Scale(v, x, y, z)
}

func (v *Vec3) ScaleSet(x, y, z float64) *Vec3 {
	ScaleSet(v, x, y, z)
	return v
}
