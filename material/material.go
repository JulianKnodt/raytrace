package material

import (
	v "raytrace/vector"
)

type Material interface {
	Color() v.Vec3
}

type ColoredMat struct {
	color v.Vec3
}

var Black = v.Vec3{0, 0, 0}
var BasicBlack = ColoredMat{Black}
var White = v.Vec3{255, 255, 255}
var BasicWhite = ColoredMat{White}
var BasicRed = ColoredMat{v.Vec3{255, 0, 0}}
var BasicGreen = ColoredMat{v.Vec3{0, 255, 0}}
var BasicBlue = ColoredMat{v.Vec3{0, 0, 255}}

func (c *ColoredMat) Color() v.Vec3 {
	if c == nil {
		return White
	}
	return c.color
}

type Placeholder struct{}

func (p Placeholder) Color() v.Vec3 {
	return v.Vec3{123, 123, 123}
}
