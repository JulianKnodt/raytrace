package octree

type BoundingBox struct {
	Center    v.Vec3
	CornerVec v.Vec3
}

func (b BoundingBox) Intersects(other BoundingBox) bool {
	return false // TODO
}
