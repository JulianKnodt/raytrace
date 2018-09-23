package shapes

import (
	v "github.com/julianknodt/vector"
	"math"
	m "raytrace/material"
	obj "raytrace/object"
)

type Sphere struct {
	center    v.Vec3
	radiusSqr float64
	*m.Material
}

func NewSphere(center v.Vec3, radius float64, mat *m.Material) *Sphere {
	return &Sphere{center, radius * radius, mat}
}

func (s Sphere) NormalAt(p v.Vec3) (v.Vec3, bool) {
	return *p.Sub(s.center), false
}

func (s Sphere) MaterialAt(v.Vec3) *m.Material {
	return s.Material
}

func (s Sphere) Intersects(r v.Ray) (a float64, shape obj.SurfaceElement) {
	center := s.center.Sub(r.Origin)
	toNormal := center.Dot(r.Direction)
	if toNormal < 0 {
		return a, nil
	}
	distSqr := center.SqrMagn() - toNormal*toNormal
	if distSqr > s.radiusSqr {
		return a, nil
	}

	interDist := math.Sqrt(s.radiusSqr - distSqr)
	t0 := toNormal - interDist
	t1 := toNormal + interDist

	if t0 < 0 {
		return t1, s
	} else {
		return t0, s
	}
}

// https://gamedev.stackexchange.com/questions/114412/how-to-get-uv-coordinates-for-sphere-cylindrical-projection
func (s Sphere) TextureCoordinates(vec v.Vec3) (u, v float64) {
	// https://github.com/fogleman/pt/blob/master/pt/sphere.go
	n := *vec.Sub(s.center).UnitSet()
	return math.Atan2(n[0], n[2])/(2*math.Pi) + 0.5,
		0.5 - n[1]*0.5
	//	return 1 - (math.Pi+math.Atan2(radial[2], radial[0]))/(2*math.Pi),
	//		(math.Atan2(radial[1], math.Hypot(radial[0], radial[2])) + math.Pi/2) / math.Pi
}

func (s Sphere) Intersects2(r v.Ray) (t float64, shape obj.SurfaceElement) {
	centerDiff := r.Origin.Sub(s.center)
	centerSqrMgn := centerDiff.SqrMagn()
	a := r.Direction.SqrMagn()
	b := 2 * centerDiff.Dot(r.Direction)
	c := centerSqrMgn - s.radiusSqr
	discrim := (b * b) - (4 * a * c)
	if discrim < 0 || a == 0 {
		return t, nil
	}

	if centerSqrMgn <= s.radiusSqr {
		return t, nil
	}

	t0 := (-b + math.Sqrt(discrim)) / (2 * a)
	t1 := (-b - math.Sqrt(discrim)) / (2 * a)
	switch {
	case t0 < 0 && t1 < 0:
		return t, nil
	case t0 < 0:
		t = t1
	case t1 < 0:
		t = t0
	default:
		// return closest point
		t = math.Min(t0, t1)
	}
	shape = s
	return
}

func (s Sphere) EmitsLight() bool {
	return s.Material.IsLighting()
}

func (s Sphere) LightOrigins() []*v.Vec3 {
	return []*v.Vec3{s.center.Copy()}
}
