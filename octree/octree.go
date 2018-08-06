package octree

import (
//	v "raytrace/vector" // only works on one of my machines lol
)

type Octree struct {
	Center [3]float64

	Parent   *Octree
	Children [8]*Octree
	Bounds   [2][3]float64

	insertionQueue []interface{} // should this be an interface
	items          []interface{}
}

const MIN_SIZE = 1
const (
	XYZ = iota
	XYNZ
	XNYZ
	XNYNZ
	NXYZ
	NXYNZ
	NXNYZ
	NXNYNZ
)

func NewOctree() *Octree {
	res := &Octree{
		Parent:   nil,
		Children: [8]*Octree{},
	}
	return res
}

func (o *Octree) Insert(item interface{}) {
	o.insertionQueue = append(o.insertionQueue, item)
}
