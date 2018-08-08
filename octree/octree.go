package octree

import (
	"math"
)

type Octree struct {
	Center [3]float64

	Parent   *Octree
	Children [8]*Octree
	bounds   NaiveBoundingBox
	Values   []OctreeItem
}

type OctreeItem interface {
	BoundingBox() NaiveBoundingBox
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
	return &Octree{
		Parent:   nil,
		Children: [8]*Octree{},
		bounds: NaiveBoundingBox{
			math.Inf(-1), math.Inf(1),
			math.Inf(-1), math.Inf(1),
			math.Inf(-1), math.Inf(1),
		},
	}
}

func (o *Octree) Insert(item ...OctreeItem) {
	o.Values = append(o.Values, item...)
}

func (o *Octree) Flatten() {
	for _, v := range o.Values {

	}
}
