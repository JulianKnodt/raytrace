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

	switch {
	case near == nil:
		return nil
	case maxDist < 0:
		panic("something behind the camera is supposed to be visible")
	}

	inter := r.Origin.Add(*r.Direction.SMul(maxDist))
	normalInter, invertable := near.NormalAt(*inter)
	material := near.MaterialAt(*inter)
	n := *v.NewRay(*inter, normalInter)
	normalInter.UnitSet()
	inter.AddSet(*normalInter.SMul(epsilon))

	// Ambient color is always regarded regardless of light
	c := new(color.Normalized)
	*c = material.Ambience()
	if s, ok := near.(mat.Sampleable); ok && material != nil && material.AmbientTexture != nil {
		u, v := s.TextureCoordinates(*inter)
		c.RGB.AddSet(texture.Sample(material.AmbientTexture, u, v).RGB)
	}

	for _, l := range s.Lights {
		origins := l.LightOrigins()

	lightSources:
		for _, origin := range origins {
			light := *v.NewRayFrom(*origin, *inter)
			facingSimilarly := normalInter.Dot(light.Direction)

			switch {
			case facingSimilarly < 0 && invertable:
				facingSimilarly = -facingSimilarly
			case facingSimilarly <= 0:
				c = c.Mix(color.Black)
				continue
			}

			// Check if it's obstructed
			for _, o := range s.Objects {
				if _, intersecting := o.Intersects(light); intersecting != nil {
					c = c.Mix(color.Black)
					continue lightSources
				}
			}

			if material.DoesReflect() {
				reflection := Direct(*v.NewRay(*inter, *r.Direction.Reflection(normalInter)), s)
				if reflection != nil {
					c = c.Mix(*reflection)
				}
			}
			c = c.MixVec(
				*l.Emitted().RGB.
					SMul(material.SurfaceBRDF().BidirectionalRadianceDistribution(n, light, r)))
		}
	}
	return c
}
