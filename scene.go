package main

import (
  "image"
  "math"
  "image/color"
)

import (
  "sync"
)

const epsilon = 1e-6

func checkIntersects(from, dir Vec3, objects []Object, lights []Light) color.Color {
  maxDist := math.Inf(1)
  var near Object
  for _, o := range objects {
    if dist, didIntersect := o.Intersects(from, dir); didIntersect {
      if dist < maxDist && dist > 0 {
        maxDist = dist
        near = o
      }
    }
  }

  if near == nil {
    return color.Black
  }

  if maxDist < 0 {
    panic("something behind the camera is supposed to be visible")
  }


  inter := Add(from, SMul(maxDist, dir))
  normalInter, invAble := near.Normal(inter)
  UnitSet(&normalInter)
  AddSet(&inter, SMul(epsilon, normalInter))

  var color Vec3
  for _, l := range lights {
    lightDir, canIllum := l.LightTo(inter) // intersection -> light
    align := Dot(normalInter, lightDir)
    if !canIllum || align <= 0 {
      if align < 0 && invAble {
        align = -align
        AddSet(&inter, SMul(-2 * epsilon, normalInter))
      } else {
        continue
      }
    }
    for _, o := range objects {
      if _, didInter := o.Intersects(inter, lightDir); didInter {
        canIllum = false
        break
      }
    }
    if canIllum {
      AddSet(&color, SMul(align, near.Color()))
    }
  }
  return ToRGBA(color)
}

func render(width, height float64, c Camera, o []Object, l []Light) image.RGBA {
  img := image.NewRGBA(image.Rect(0,0, int(width), int(height)))
  var invWidth float64 = 1.0/width
  var invHeight float64 = 1.0/height
  aspectRatio := width * invHeight
  angle := math.Tan(math.Pi * 0.5 * c.FOV()/180)
  for y := 0.0; y < height; y++ {
    for x := 0.0; x < width; x++ {
      xDir := (2 * ((x + 0.5) * invWidth) - 1) * angle * aspectRatio
      yDir := (1 - 2  * ((y + 0.5) * invHeight)) * angle
      direction := Unit(Sub(Vec3{xDir, yDir, -1}, c.Location()))
      img.Set(int(x),int(y), checkIntersects(Origin, direction, o, l))
    }
  }
  return *img
}
