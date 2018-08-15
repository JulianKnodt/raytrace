package octree

import (
	"raytrace/bounding"
)

func CreateFrom(from interface {
	BoundingBox() bounding.AxisAlignedBoundingBox
	Children() []OctreeItem
}) *Octree {
	result := NewEmptyOctree(from.BoundingBox())

	result.Insert(from.Children()...)
	result.Flatten()
	return result
}
