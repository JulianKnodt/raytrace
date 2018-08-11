package vector

// Represents a ray starting from origin
// going in direction dir
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func NewRay(origin, dir Vec3) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: Unit(dir),
	}
}
