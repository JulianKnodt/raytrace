package object

import (
	m "raytrace/material"
	v "raytrace/vector"
)

/*
 Intersects: an object should have the
 ability to check whether it intersects with
 a vector

 Normal: an object should be able to determine a normal on its surface

 Color: an object should have a color
*/
type Object interface {
	Intersects(origin, dir v.Vec3) (float64, SurfaceElement)
}

type SurfaceElement interface {
	Normal(to v.Vec3) (dir v.Vec3, invAble bool)
	// invable relates to whether or not the normal can be flipped or not
	// any 2d shape should be invertible
	Mat() m.Material
}
