package obj

import (
	vec "github.com/julianknodt/vector"
	itl "raytrace/indexedTriangleList"
)

func (o Obj) IndexedTriangleList() *itl.IndexedTriangleList {
	t := itl.NewIndexedTriangleList()
	for i := 1; i < len(o.V); i++ {
		t.AddVertex(vec.Vec3(o.V[i]))
	}

	for i := 1; i < len(o.Vt); i++ {
		t.AddTexture(vec.Vec3(o.Vt[i]))
	}

	for i := 1; i < len(o.Vn); i++ {
		t.AddNormal(vec.Vec3(o.Vn[i]))
	}

	for _, face := range o.F {
		t.AddFace(
			face.Vertices(),
			face.Normals(),
			face.Textures(),
			o.MTLLib[face.Material].Material(),
		)
	}

	return t
}
