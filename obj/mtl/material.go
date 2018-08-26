package mtl

import (
	"log"
	"raytrace/color"
	"raytrace/material"
)

func (m MTL) Material() *material.Material {
	out := &material.Material{}
	out.Ambient = color.FromNormalized(m.Ka[0], m.Ka[1], m.Ka[2], 1)
	out.Emissive = color.FromNormalized(m.Ke[0], m.Ke[1], m.Ke[2], 1)
	out.Diffuse = color.FromNormalized(m.Kd[0], m.Kd[1], m.Kd[2], 1)
	if img, err := m.Map_Kd.Load(m.fileName); err == nil {
		out.DiffuseTexture = img
	} else {
		log.Println(err)
	}

	if img, err := m.Map_Bump.Load(m.fileName); err == nil {
		out.BumpTexture = img
	} else {
		log.Println(err)
	}
	return out
}
