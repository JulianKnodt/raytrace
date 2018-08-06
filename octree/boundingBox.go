package octree

import (
	v "raytrace/vector"
)

type BoundingBox struct {
	Center    v.Vec3
	CornerVec v.Vec3
}

func (b BoundingBox) Intersects(other BoundingBox) bool {
	return false // TODO
}
