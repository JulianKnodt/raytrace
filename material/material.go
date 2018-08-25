package material

// Represents the material of a surface element
type Material interface {
	Emitted() [3]float64
}

type ColoredMat struct {
	color [3]float64
}

var Black = [3]float64{0, 0, 0}
var BasicBlack = ColoredMat{Black}
var White = [3]float64{255, 255, 255}
var BasicWhite = ColoredMat{White}
var BasicRed = ColoredMat{[3]float64{255, 0, 0}}
var BasicGreen = ColoredMat{[3]float64{0, 255, 0}}
var BasicBlue = ColoredMat{[3]float64{0, 0, 255}}

func (c *ColoredMat) Emitted() [3]float64 {
	if c == nil {
		return White
	}
	return c.color
}

type Placeholder struct{}

func (p Placeholder) Emitted() [3]float64 {
	return [3]float64{123, 123, 123}
}
