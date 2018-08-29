package material

import (
	"raytrace/color"
	texture "raytrace/material/texture"
)

// http://blog.lexique-du-net.com/index.php?post/2009/07/24/AmbientDiffuseEmissive-and-specular-colorSome-examples
// https://en.wikipedia.org/wiki/Shading#Ambient_lighting
// https://computergraphics.stackexchange.com/questions/375/what-is-ambient-lighting

// Represents the material of a surface element
type Material struct {
	Ambient  color.Normalized
	Diffuse  color.Normalized
	Emissive color.Normalized
	// Specular will be ignored for now...

	BumpTexture texture.Texture

	DiffuseTexture texture.Texture
}

func (m *Material) Emitted() color.Normalized {
	if m == nil {
		return color.DefaultColor
	}
	return m.Emissive
}

func (m *Material) Ambience() color.Normalized {
	if m == nil {
		return color.DefaultColor
	}
	return m.Ambient
}

func (m *Material) Diffusive() color.Normalized {
	if m == nil {
		return color.DefaultColor
	}
	return m.Diffuse
}
