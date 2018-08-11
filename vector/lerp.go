package vector

func Lerp(from, to Vec3, proportion float64) Vec3 {
	switch {
	case proportion >= 1:
		return to
	case proportion <= 0:
		return from
	}
	return Add(from, SMul(proportion, Sub(to, from)))
}
