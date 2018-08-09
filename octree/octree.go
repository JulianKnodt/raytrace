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

func NewOctree(bounds NaiveBoundingBox) *Octree {
	return &Octree{
		Parent:   nil,
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

	center := o.Region.Center()
	if o.Children[XYZ] == nil {
		o.Children[XYZ] = &Octree{
			Parent:   o,
			Children: [8]*Octree{},
			Region:   NaiveBoundingBox{center[0], o.Region.XX, center[1], o.Region.YY, center[2], o.Region.ZZ},
		}
	}
}
