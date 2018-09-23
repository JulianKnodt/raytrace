package material

import (
	"raytrace/color"
)

func WhiteLightMaterial() *Material {
	return &Material{
		Ambient:      color.Blank,
		Diffuse:      color.Blank,
		Emissive:     color.White,
		Transparency: 0,
	}
}
