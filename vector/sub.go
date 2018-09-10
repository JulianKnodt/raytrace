package vector

func Sub(a, b Vec3) (r Vec3) {
	r[0] = a[0] - b[0]
	r[1] = a[1] - b[1]
	r[2] = a[2] - b[2]
	return
}

func (a Vec3) Sub(b Vec3) (r Vec3) {
	r[0] = a[0] - b[0]
	r[1] = a[1] - b[1]
	r[2] = a[2] - b[2]
	return
}

func (a *Vec3) SubSet(b Vec3) *Vec3 {
	a[0] = a[0] - b[0]
	a[1] = a[1] - b[1]
	a[2] = a[2] - b[2]
	return a
}
