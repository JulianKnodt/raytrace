package indexedTriangleList

import (
	"raytrace/octree"
)

func (itl IndexedTriangleList) Children() []octree.OctreeItem {
	size := itl.Size()
	out := make([]octree.OctreeItem, size)
	for i := 0; i < size; i++ {
		out[i] = itl.GetTriangle(i)
	}
	return out
}
