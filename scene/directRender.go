package scene

import (
	"image/color"
	"math"

	"raytrace/object"
	v "raytrace/vector"
)

func Direct(from, dir v.Vec3, s Scene) color.Color {
	maxDist := math.Inf(1)
	var near object.SurfaceElement
	for _, o := range s.Objects {
		if dist, intersecting := o.Intersects(from, dir); intersecting != nil {
			if dist < maxDist && dist > 0 {
				maxDist = dist
				near = intersecting
			}
		}
	}

	if near == nil {
		return color.Black
	}

	if maxDist < 0 {
		panic("something behind the camera is supposed to be visible")
	}

	inter := v.Add(from, v.SMul(maxDist, dir))
	normalInter, invAble := near.Normal()
	v.UnitSet(&normalInter)
	v.AddSet(&inter, v.SMul(epsilon, normalInter))

	var color v.Vec3
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
			if _, intersecting := o.Intersects(inter, lightDir); intersecting != nil {
				canIllum = false
				break
			}
		}
		if canIllum {
			v.AddSet(&color, v.SMul(align, near.Mat().Emitted()))
		}
	}
	return v.ToRGBA(color)
}
