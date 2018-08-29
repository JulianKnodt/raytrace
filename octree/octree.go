package octree

import (
	"raytrace/bounding"
	"raytrace/object"
)

type Octree struct {
	Parent            *Octree
	Children          [8]*Octree
	processedValues   []OctreeItem
	UnprocessedValues []OctreeItem
	Region            bounding.Box
}

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

func NewEmptyOctree(bounds bounding.Box) *Octree {
	return &Octree{
		Parent:            nil,
		Children:          [8]*Octree{},
		Region:            bounds,
		UnprocessedValues: make([]OctreeItem, 0),
		processedValues:   make([]OctreeItem, 0),
	}
}

func newChildOctree(parent *Octree, bounds bounding.Box) *Octree {
	return &Octree{
		Parent:            parent,
		Children:          [8]*Octree{},
		Region:            bounds,
		UnprocessedValues: make([]OctreeItem, 0),
		processedValues:   make([]OctreeItem, 0),
	}
}

type OctreeItem interface {
	Box() bounding.Box
	object.Object
}

func (o *Octree) Insert(items ...OctreeItem) {
	o.UnprocessedValues = append(o.UnprocessedValues, items...)
}

func (o *Octree) Len() int {
	if o == nil {
		return -1
	}
	return len(o.processedValues) + len(o.UnprocessedValues)
}

func (o *Octree) Flatten() {
	if o == nil || len(o.UnprocessedValues) == 0 {
		return
	}

	// Creating all the children

	center := o.Region.Center()
	cX, cY, cZ := center[0], center[1], center[2]
	if o.Children[XYZ] == nil {
		o.Children[XYZ] = newChildOctree(o, *bounding.NewBox(
			cX, o.Region.Max[0],
			cY, o.Region.Max[1],
			cZ, o.Region.Max[2],
		))
	}

	if o.Children[XYNZ] == nil {
		o.Children[XYNZ] = newChildOctree(o, *bounding.NewBox(
			cX, o.Region.Max[0],
			cY, o.Region.Max[1],
			o.Region.Min[2], cZ,
		))
	}

	if o.Children[XNYZ] == nil {
		o.Children[XNYZ] = newChildOctree(o, *bounding.NewBox(
			cX, o.Region.Max[0],
			o.Region.Min[1], cY,
			cZ, o.Region.Max[2],
		))
	}

	if o.Children[XNYNZ] == nil {
		o.Children[XNYNZ] = newChildOctree(o, *bounding.NewBox(
			cX, o.Region.Max[0],
			o.Region.Min[1], cY,
			o.Region.Min[2], cZ,
		))
	}

	if o.Children[NXYZ] == nil {
		o.Children[NXYZ] = newChildOctree(o, *bounding.NewBox(
			o.Region.Min[0], cX,
			cY, o.Region.Max[1],
			cZ, o.Region.Max[2],
		))
	}

	if o.Children[NXYNZ] == nil {
		o.Children[NXYNZ] = newChildOctree(o, *bounding.NewBox(
			o.Region.Min[0], cX,
			cY, o.Region.Max[1],
			o.Region.Min[2], cZ,
		))
	}

	if o.Children[NXNYZ] == nil {
		o.Children[NXNYZ] = newChildOctree(o, *bounding.NewBox(
			o.Region.Min[0], cX,
			o.Region.Min[1], cY,
			cZ, o.Region.Max[2],
		))
	}

	if o.Children[NXNYNZ] == nil {
		o.Children[NXNYNZ] = newChildOctree(o, *bounding.NewBox(
			o.Region.Min[0], cX,
			o.Region.Min[1], cY,
			o.Region.Min[2], cZ,
		))
	}

	for _, item := range o.UnprocessedValues {
		var intersected *Octree

	CheckSubregions:
		for _, sub := range o.Children {
			contains := sub.Region.Contains(item.Box())
			if contains {
				switch {
				case intersected == nil:
					intersected = sub
				default:
					intersected = nil
					break CheckSubregions
				}
			}
		}

		if intersected == nil {
			o.processedValues = append(o.processedValues, item)
		} else {
			intersected.Insert(item)
		}
	}

	o.UnprocessedValues = make([]OctreeItem, 0)

	for _, sub := range o.Children {
		sub.Flatten()
	}
}
