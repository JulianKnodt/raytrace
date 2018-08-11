package vector

func Add(a, b Vec3) (r Vec3) {
	r[0] = a[0] + b[0]
	r[1] = a[1] + b[1]
	r[2] = a[2] + b[2]
	return
}

func AddSet(dest *Vec3, src Vec3) {
	dest[0] = dest[0] + src[0]
	dest[1] = dest[1] + src[1]
	dest[2] = dest[2] + src[2]
}

func (v Vec3) Add(o Vec3) Vec3 {
	return Add(v, o)
}

func (v *Vec3) AddSet(o Vec3) *Vec3 {
	AddSet(v, o)
	return v
}