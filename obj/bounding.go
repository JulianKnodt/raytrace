package obj

import (
	"math"
	"raytrace/bounding"
	"raytrace/octree"
	"raytrace/shapes"
)

func (o Obj) BoundingBox() bounding.AxisAlignedBoundingBox {
	Xx, Yy, Zz := math.Inf(1), math.Inf(1), math.Inf(1)
	XX, YY, ZZ := math.Inf(-1), math.Inf(-1), math.Inf(-1)
	for _, v := range o.V {
		if v[0] < Xx {
			Xx = v[0]
		} else if v[0] > XX {
			XX = v[0]
		}

		if v[1] < Yy {
			Yy = v[1]
		} else if v[1] > YY {
			YY = v[1]
		}

		if v[2] < Zz {
			Zz = v[2]
		} else if v[2] > ZZ {
			ZZ = v[2]
		}
	}
	return bounding.AxisAlignedBoundingBox{
		Xx: Xx, XX: XX, Yy: Yy, YY: YY, Zz: Zz, ZZ: ZZ,
	}
}

func (o Obj) Triangles() []shapes.Triangle {
	result := make([]shapes.Triangle, 0, len(o.F))
	for i := 0; i < len(o.F); i++ {
		result = append(result, shapes.ToTriangles(o.ShapeN(i), o.TextureN(i))...)
	}
	return result
}

func (o Obj) Children() []octree.OctreeItem {
	triangles := o.Triangles()
	result := make([]octree.OctreeItem, 0, len(triangles))
	for i, v := range triangles {
		result[i] = v
	}
	return result
}
