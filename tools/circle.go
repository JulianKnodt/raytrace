package main

import (
	"fmt"
	"io"
	"math"
	v "raytrace/vector"
)

// outputs a Cylinder to w
// with given height
// in off file format
// with base at height 0
func Out(w io.Writer, height, radius float64, num int) {
	// x y z boi
	base := Circle(radius, num)
	top := v.Shift(base, 0, 0, 10)
	numFaces := 2 + num
	writeHeader(w, 2*num, numFaces, 13) // num edges is whatever
	for _, point := range base {
		w.Write([]byte(fmt.Sprintf("%f %f %f\n", point.X(), point.Y(), point.Z())))
	}
	for _, point := range top {
		w.Write([]byte(fmt.Sprintf("%f %f %f\n", point.X(), point.Y(), point.Z())))
	}

	writeFaces(w, 0, len(base), num)
}

func writeHeader(w io.Writer, numVertices, numFaces, numEdges int) {
	w.Write([]byte("OFF\n\n"))
	w.Write([]byte(fmt.Sprintf("%d %d %d\n\n", numVertices, numFaces, numEdges)))
}

func writeFaces(w io.Writer, startA, startB, num int) {
	for i := 0; i < num-1; i++ {
		w.Write([]byte(fmt.Sprintf("4 %d %d %d %d\n",
			startA+i,
			startB+i,
			startA+i+1,
			startB+i+1,
		)))
	}
	w.Write([]byte(fmt.Sprintf("4 %d %d %d %d\n", startA, startB, startA+num-1, startB+num-1)))
}

func Circle(radius float64, num int) []v.Vec3 {
	sliceRads := math.Pi * 2 / float64(num)
	out := make([]v.Vec3, 0)
	for i := 0; i < num; i++ {
		rads := float64(i) * sliceRads
		out = append(out, v.Vec3{math.Cos(rads) * radius, math.Sin(rads) * radius, 0})
	}
	return out
}
