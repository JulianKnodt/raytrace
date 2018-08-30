package scene

import (
	"image/color"
	"math"

	mat "raytrace/material"
	texture "raytrace/material/texture"
	"raytrace/object"
	v "raytrace/vector"
)

func Direct(r v.Ray, s Scene) color.Color {
	maxDist := math.Inf(1)
	var near object.SurfaceElement
	for _, o := range s.Objects {
		if dist, intersecting := o.Intersects(r); intersecting != nil {
			if dist < maxDist && dist > 0 {
				maxDist = dist
				near = intersecting
			}
		}
	}

	if near == nil {
		return nil
	}

	if maxDist < 0 {
		panic("something behind the camera is supposed to be visible")
	}

	inter := v.Add(r.Origin, v.SMul(maxDist, r.Direction))
	normalInter, invAble := near.NormalAt(inter)
	material := near.MaterialAt(inter)
	v.UnitSet(&normalInter)
	v.AddSet(&inter, v.SMul(epsilon, normalInter))

	// Ambient color is regardless
	color := material.Ambience()
	if s, ok := near.(mat.Sampleable); ok && material != nil && material.AmbientTexture != nil {
		u, v := s.TextureCoordinates(inter)
		color.RGB = color.RGB.Add(texture.Sample(material.AmbientTexture, u, v).RGB)
	}

	for _, l := range s.Lights {
		lightDir, canIllum := l.LightTo(inter) // intersection -> light
		align := v.Dot(normalInter, lightDir)
		if !canIllum || align <= 0 {
			if align < 0 && invAble {
				align = -align
				v.AddSet(&inter, v.SMul(-2*epsilon, normalInter))
			} else {
				continue
			}
		}
		for _, o := range s.Objects {
			if _, intersecting := o.Intersects(*v.NewRay(inter, lightDir)); intersecting != nil {
				canIllum = false
				break
			}
		}
		if canIllum {
			color.RGB = color.RGB.Add(material.Diffusive().RGB.SMul(align))
		}
	}
	return color.ToImageColor()
}
