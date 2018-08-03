package obj

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	v "raytrace/vector"
)

// Direct OBJ representation
type Obj struct {
	// Geometric Vertices
	V map[int]v.Vec3
	// Polygonal Face Element
	F [][][]int
	// Line Element
	L [][]int
}

func Decode(r io.Reader) (*Obj, error) {
	scanner := bufio.NewScanner(r)
	o := &Obj{
		V: make(map[int]v.Vec3),
		F: [][][]int{},
	}
	for scanner.Scan() {
		if err := o.add(scanner.Text()); err != nil {
			return o, err
		}
	}
	return o, nil
}

func (o *Obj) add(s string) error {
	parts := strings.SplitN(s, " ", 2)
	switch parts[0] {
	case "", "#":
	case "v":
		var x, y, z float64
		_, err := fmt.Sscanf(parts[1], "%f %f %f", &x, &y, &z)
		if err != nil {
			return err
		}
		o.V[len(o.V)+1] = v.Vec3{x, y, z}
	case "f":
		face := [][]int{}
		for _, coords := range strings.Fields(parts[1]) {
			switch strings.Count(coords, "/") {
			case 0:
				var vertNumber int
				_, err := fmt.Sscanf(coords, "%d", &vertNumber)
				if err != nil {
					return err
				}
				face = append(face, []int{vertNumber})
			case 1:
				var a, b int
				_, err := fmt.Sscanf(coords, "%d/%d", &a, &b)
				if err != nil {
					return err
				}
				face = append(face, []int{a, b})
			case 2:
				var a, b, c int
				_, err := fmt.Sscanf(coords, "%d/%d/%d", &a, &b, &c)
				if err != nil {
					return err
				}
				face = append(face, []int{a, b, c})
			}
		}
		o.F = append(o.F, face)
	}
	return nil
}

func Encode(w io.Writer, o Obj) error {
	return nil
}
