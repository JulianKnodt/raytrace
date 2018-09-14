package indexedTriangleList

import (
	v "github.com/julianknodt/vector"
	"math"
	"raytrace/material"
	obj "raytrace/object"
	"raytrace/shapes"
)

type FaceInfo struct {
	Vertices [3]int
	Normals  *[3]int
	Textures *[3]int
	Material *material.Material
}

type IndexedTriangleList struct {
	Vertices        []v.Vec3
	NormalVertices  []v.Vec3
	TextureVertices []v.Vec3

	Faces []FaceInfo
}

func (t IndexedTriangleList) GetTriangle(nth int) *shapes.Triangle {
	face := t.Faces[nth]
	out := shapes.NewTriangle(
		t.Vertices[face.Vertices[0]],
		t.Vertices[face.Vertices[1]],
		t.Vertices[face.Vertices[2]],
		face.Material,
	)

	if face.Normals != nil {
		out.SetNormals(
			t.NormalVertices[face.Normals[0]],
			t.NormalVertices[face.Normals[1]],
			t.NormalVertices[face.Normals[2]],
		)
	}

	if face.Textures != nil {
		out.SetTextures(
			t.TextureVertices[face.Textures[0]],
			t.TextureVertices[face.Textures[1]],
			t.TextureVertices[face.Textures[2]],
		)
	}

	return out
}

func (t IndexedTriangleList) Size() int {
	return len(t.Faces)
}

func NewIndexedTriangleList() *IndexedTriangleList {
	return &IndexedTriangleList{
		make([]v.Vec3, 0),
		make([]v.Vec3, 0),
		make([]v.Vec3, 0),
		make([]FaceInfo, 0),
	}
}

func (i *IndexedTriangleList) AddVertex(vec v.Vec3) {
	i.Vertices = append(i.Vertices, vec)
}

func (i *IndexedTriangleList) AddNormal(vec v.Vec3) {
	i.NormalVertices = append(i.NormalVertices, vec)
}

func (i *IndexedTriangleList) AddTexture(vec v.Vec3) {
	i.TextureVertices = append(i.TextureVertices, vec)
}

func (itl *IndexedTriangleList) AddFace(indeces []int,
	normals *[]int,
	textures *[]int,
	mat *material.Material) {
	if len(indeces) < 3 {
		panic("Cannot add verteces of length less than 3")
	}

	for i := 0; i < len(indeces)-2; i++ {
		face := FaceInfo{}
		face.Vertices = [3]int{indeces[i], indeces[i+1], indeces[i+2]}
		face.Material = mat
		if normals != nil {
			face.Normals = &[3]int{(*normals)[i], (*normals)[i+1], (*normals)[i+2]}
		}
		if textures != nil {
			face.Textures = &[3]int{(*textures)[i], (*textures)[i+1], (*textures)[i+2]}
		}
		itl.Faces = append(itl.Faces, face)
	}
}

func (t IndexedTriangleList) Intersects(r v.Ray) (float64, obj.SurfaceElement) {
	pMax := math.Inf(1)
	var surface shapes.Triangle
	size := t.Size()
	for i := 0; i < size; i++ {
		curr := t.GetTriangle(i)
		if p, hit := curr.Intersects(r); hit != nil {
			if p < pMax {
				pMax = p
				surface = *curr
			}
		}
	}
	return pMax, surface
}
