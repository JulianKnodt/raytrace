package octree

func (o Octree) Iterate(callback func(Octree) bool) {
	if callback(o) {
		for _, sub := range o.Children {
			if sub != nil {
				sub.Iterate(callback)
			}
		}
	}
}
