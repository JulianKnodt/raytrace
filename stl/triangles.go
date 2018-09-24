package stl

import (
	"raytrace/material"
	"raytrace/shapes"
)

func ToTriangles(ts []Triangle, mat *material.Material) []shapes.Triangle {
	out := make([]shapes.Triangle, len(ts))
	for i, t := range ts {
		newT := shapes.NewTriangle(t.V1, t.V2, t.V3, mat)
		out[i] = *newT
	}
	return out
}
