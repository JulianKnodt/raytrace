package material

import (
	"raytrace/color"
	texture "raytrace/material/texture"
)

// http://blog.lexique-du-net.com/index.php?post/2009/07/24/AmbientDiffuseEmissive-and-specular-colorSome-examples

// Represents the material of a surface element
type Material struct {
	Ambient  color.Normalized
	Diffuse  color.Normalized
	Emissive color.Normalized
	// Specular will be ignored for now...

	BumpTexture texture.Texture

	DiffuseTexture texture.Texture
}
