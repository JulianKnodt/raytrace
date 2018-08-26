package octree

func CreateFrom(from interface {
	OctreeItem
	Children() []OctreeItem
}) *Octree {
	result := NewEmptyOctree(from.Box())

	result.Insert(from.Children()...)
	result.Flatten()
	return result
}
