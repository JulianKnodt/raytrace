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
	// Specular color.Normalized will be ignored for now...

	// [0, 1], 0 fully transparent or 1 for fully opaque
	Transparency float64

	// Whether or not this material reflects light
	// Currently only reflects light once
	Reflect bool

	BumpTexture texture.Texture

	AmbientTexture texture.Texture
	DiffuseTexture texture.Texture

	RenderType Scatter
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

func (m *Material) DoesReflect() bool {
	if m == nil {
		return false
	}
	return m.Reflect
}

func (m *Material) IsLighting() bool {
	switch {
	case m == nil:

		// Checking if there is some Emission of light
	case m.Emissive.A > 0:
		return true

	}
	return false
}

var DefaultSurfaceBRDF = Lambertian{1}

func (m *Material) SurfaceBRDF() Scatter {
	switch {
	case m == nil, m.RenderType == nil:
		return DefaultSurfaceBRDF
	default:
		return m.RenderType
	}
}
