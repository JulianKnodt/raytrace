package octree

type Octree struct {
	Parent   *Octree
	Children [8]*Octree
	Values   []OctreeItem
	Region   NaiveBoundingBox
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

func NewEmptyOctree(bounds NaiveBoundingBox) *Octree {
	return &Octree{
		Parent:   nil,
		Children: [8]*Octree{},
		Region:   bounds,
		Values:   make([]OctreeItem, 0),
	}
}

func newChildOctree(parent *Octree, bounds NaiveBoundingBox) *Octree {
	return &Octree{
		Parent:   parent,
		Children: [8]*Octree{},
		Region:   bounds,
		Values:   make([]OctreeItem, 0),
	}
}

type OctreeItem interface {
	Box() NaiveBoundingBox
}

func (o *Octree) Insert(items ...OctreeItem) {
	o.Values = append(o.Values, items...)
	o.Flatten(8)
}

func (o *Octree) Flatten(allowedAmt int) {
	if len(o.Values) <= allowedAmt {
		return
	}

	// Creating all the children

	center := o.Region.Center()
	cX, cY, cZ := center[0], center[1], center[2]
	if o.Children[XYZ] == nil {
		o.Children[XYZ] = newChildOctree(o, NaiveBoundingBox{
			cX, o.Region.XX,
			cY, o.Region.YY,
			cZ, o.Region.ZZ,
		})
	}

	if o.Children[XYNZ] == nil {
		o.Children[XYNZ] = newChildOctree(o, NaiveBoundingBox{
			cX, o.Region.XX,
			cY, o.Region.YY,
			o.Region.Zz, cZ,
		})
	}

	if o.Children[XNYZ] == nil {
		o.Children[XNYZ] = newChildOctree(o, NaiveBoundingBox{
			cX, o.Region.XX,
			o.Region.Yy, cY,
			cZ, o.Region.ZZ,
		})
	}

	if o.Children[XNYNZ] == nil {
		o.Children[XNYNZ] = newChildOctree(o, NaiveBoundingBox{
			cX, o.Region.XX,
			o.Region.Yy, cY,
			o.Region.Zz, cZ,
		})
	}

	if o.Children[NXYZ] == nil {
		o.Children[NXYZ] = newChildOctree(o, NaiveBoundingBox{
			o.Region.Xx, cX,
			cY, o.Region.YY,
			cZ, o.Region.ZZ,
		})
	}

	if o.Children[NXYNZ] == nil {
		o.Children[NXYNZ] = newChildOctree(o, NaiveBoundingBox{
			o.Region.Xx, cX,
			cY, o.Region.YY,
			o.Region.Zz, cZ,
		})
	}

	if o.Children[NXNYZ] == nil {
		o.Children[NXNYZ] = newChildOctree(o, NaiveBoundingBox{
			o.Region.Xx, cX,
			o.Region.Yy, cY,
			cZ, o.Region.ZZ,
		})
	}

	if o.Children[NXNYNZ] == nil {
		o.Children[NXNYNZ] = newChildOctree(o, NaiveBoundingBox{
			o.Region.Xx, cX,
			o.Region.Yy, cY,
			o.Region.Zz, cZ,
		})
	}
}
