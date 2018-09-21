package scene

import (
	"math"

	v "github.com/julianknodt/vector"
	"raytrace/color"
	mat "raytrace/material"
	texture "raytrace/material/texture"
	"raytrace/object"
)

func Direct(r v.Ray, s Scene) *color.Normalized {
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

	inter := r.Origin.Add(*r.Direction.SMul(maxDist))
	normalInter, _ := near.NormalAt(*inter)
	material := near.MaterialAt(*inter)
	v.UnitSet(&normalInter)
	inter.AddSet(*normalInter.SMul(epsilon))
	bounce := *v.NewRay(*inter, normalInter)

	// Ambient color is always regarded regardless of light
	c := new(color.Normalized)
	*c = material.Ambience()
	if s, ok := near.(mat.Sampleable); ok && material != nil && material.AmbientTexture != nil {
		u, v := s.TextureCoordinates(*inter)
		c.RGB.AddSet(texture.Sample(material.AmbientTexture, u, v).RGB)
	}

	for _, l := range s.Lights {
		t, surfel := l.Intersects(bounce)
		if surfel == nil {
			continue
		}
		mat := surfel.MaterialAt(*bounce.At(t))
		c.Mix(mat.Emitted())
	}
	return c
}
