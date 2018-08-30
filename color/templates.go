package color

import (
	v "raytrace/vector"
)

var DefaultColor = Normalized{
	v.Vec3{0.5, 0.5, 0.5},
	maxUint8,
}

var Blank = Normalized{
	v.Vec3{0, 0, 0},
	maxUint8,
}
