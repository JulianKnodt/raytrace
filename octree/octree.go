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
	Region            bounding.AxisAlignedBoundingBox
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

var zero = [3]float64{0, 0, 0}

func NewEmptyOctree(bounds bounding.AxisAlignedBoundingBox) *Octree {
	return &Octree{
		Parent:            nil,
		Children:          [8]*Octree{},
		Region:            bounds,
		UnprocessedValues: make([]OctreeItem, 0),
	}
}

func newChildOctree(parent *Octree, bounds bounding.AxisAlignedBoundingBox) *Octree {
	return &Octree{
		Parent:            parent,
		Children:          [8]*Octree{},
		Region:            bounds,
		UnprocessedValues: make([]OctreeItem, 0),
		processedValues:   make([]OctreeItem, 0),
	}
}

type OctreeItem interface {
	Box() bounding.AxisAlignedBoundingBox
	object.Object
}

func (o *Octree) Insert(items ...OctreeItem) {
	o.UnprocessedValues = append(o.UnprocessedValues, items...)
}

func (o *Octree) Flatten() {
	if o == nil || len(o.UnprocessedValues) == 0 {
		return
	}

	// Creating all the children

	center := o.Region.Center()
	cX, cY, cZ := center[0], center[1], center[2]
	if o.Children[XYZ] == nil {
		o.Children[XYZ] = newChildOctree(o, bounding.AxisAlignedBoundingBox{
			Xx: cX, XX: o.Region.XX,
			Yy: cY, YY: o.Region.YY,
			Zz: cZ, ZZ: o.Region.ZZ,
		})
	}

	if o.Children[XYNZ] == nil {
		o.Children[XYNZ] = newChildOctree(o, bounding.AxisAlignedBoundingBox{
			Xx: cX, XX: o.Region.XX,
			Yy: cY, YY: o.Region.YY,
			Zz: o.Region.Zz, ZZ: cZ,
		})
	}

	if o.Children[XNYZ] == nil {
		o.Children[XNYZ] = newChildOctree(o, bounding.AxisAlignedBoundingBox{
			Xx: cX, XX: o.Region.XX,
			Yy: o.Region.Yy, YY: cY,
			Zz: cZ, ZZ: o.Region.ZZ,
		})
	}

	if o.Children[XNYNZ] == nil {
		o.Children[XNYNZ] = newChildOctree(o, bounding.AxisAlignedBoundingBox{
			Xx: cX, XX: o.Region.XX,
			Yy: o.Region.Yy, YY: cY,
			Zz: o.Region.Zz, ZZ: cZ,
		})
	}

	if o.Children[NXYZ] == nil {
		o.Children[NXYZ] = newChildOctree(o, bounding.AxisAlignedBoundingBox{
			Xx: o.Region.Xx, XX: cX,
			Yy: cY, YY: o.Region.YY,
			Zz: cZ, ZZ: o.Region.ZZ,
		})
	}

	if o.Children[NXYNZ] == nil {
		o.Children[NXYNZ] = newChildOctree(o, bounding.AxisAlignedBoundingBox{
			Xx: o.Region.Xx, XX: cX,
			Yy: cY, YY: o.Region.YY,
			Zz: o.Region.Zz, ZZ: cZ,
		})
	}

	if o.Children[NXNYZ] == nil {
		o.Children[NXNYZ] = newChildOctree(o, bounding.AxisAlignedBoundingBox{
			Xx: o.Region.Xx, XX: cX,
			Yy: o.Region.Yy, YY: cY,
			Zz: cZ, ZZ: o.Region.ZZ,
		})
	}

	if o.Children[NXNYNZ] == nil {
		o.Children[NXNYNZ] = newChildOctree(o, bounding.AxisAlignedBoundingBox{
			Xx: o.Region.Xx, XX: cX,
			Yy: o.Region.Yy, YY: cY,
			Zz: o.Region.Zz, ZZ: cZ,
		})
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
