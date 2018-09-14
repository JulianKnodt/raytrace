package color

import (
	v "raytrace/vector"
)

var DefaultColor = Normalized{
	v.Vec3{0.5, 0.5, 0.5},
	1,
}

var Blank = Normalized{
	v.Vec3{0, 0, 0},
	1,
}

var White = Normalized{
	v.Vec3{1, 1, 1},
	1,
}
