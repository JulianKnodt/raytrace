package vector

func Lerp(from, to Vec3, proportion float64) Vec3 {
	return Add(from, SMul(proportion, Sub(to, from)))
}

func (from Vec3) Lerp(to Vec3, proportion float64) Vec3 {
	return from.Add(to.Sub(from).SMul(proportion))
}
