package vector

type Polygon struct {
	points []Vec3
}

type errorNotCoplanar string

func (e errorNotCoplanar) Error() string { return string(e) }

const ErrorNotCoplanar = errorNotCoplanar("Points in polygon not coplanar")

func NewPolygon(from []Vec3) (*Polygon, error) {
	if !Coplanar(from) {
		return nil, ErrorNotCoplanar
	}
	return &Polygon{from}, nil
}

func (p Polygon) Normal(to Vec3) (dir Vec3, invAble bool) {
	return normalNoCheck(p.points[0], p.points[1], p.points[2])
}

func (p Polygon) Color() Vec3 {
	return Vec3{200, 200, 200}
}
