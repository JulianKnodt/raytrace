package obj

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Direct OBJ representation
type Obj struct {
	// Geometric Vertices
	V  [][4]float64
	Vt [][3]float64
	Vn [][3]float64
	Vp [][3]float64
	// Polygonal Face Element
	F [][][]int
	// Line Element
	L [][]int
}

type State struct {
	G      string
	O      string
	UseMTL string
	// MTLLib string // this loads materials from file into memory
}

func Decode(r io.Reader) (*Obj, error) {
	scanner := bufio.NewScanner(r)
	o := &Obj{
		V:  make([][4]float64, 1),
		Vt: make([][3]float64, 1),
		Vn: make([][3]float64, 1),
		F:  [][][]int{},
	}
	for scanner.Scan() {
		if err := o.add(scanner.Text()); err != nil {
			return o, err
		}
	}
	return o, nil
}

func (o *Obj) add(s string) (err error) {
	parts := strings.SplitN(s, " ", 2)
	if len(parts) == 2 {
		parts[1] = strings.TrimSpace(parts[1])
	}
	switch parts[0] {
	case "", "#":
	case "g":
		// TODO group name
	case "o":
		// TODO object name
	case "usemtl":
		// TODO something
	case "mtllib":
		// TODO something
	case "v":
		var x, y, z float64
		w := 1.0 // w is optional and defaults to 0, it is the weight of a vertex
		switch s := strings.Fields(parts[1]); len(s) {
		case 3:
			_, err = fmt.Sscanf(s[0], "%f", &x)
			_, err = fmt.Sscanf(s[1], "%f", &y)
			_, err = fmt.Sscanf(s[2], "%f", &z)
		case 4:
			_, err = fmt.Sscanf(s[0], "%f", &x)
			_, err = fmt.Sscanf(s[1], "%f", &y)
			_, err = fmt.Sscanf(s[2], "%f", &z)
			_, err = fmt.Sscanf(s[3], "%f", &w)
		default:
			err = fmt.Errorf("Cannot handle %s in obj file", parts[1])
		}
		if err != nil {
			return
		}

		o.V = append(o.V, [4]float64{x, y, z, w})
	case "vt":
		var u, vv, w float64
		_, err = fmt.Sscanf(parts[1], "%f %f %f", &u, &vv, &w)
		if err != nil {
			return
		}
		o.Vt = append(o.Vt, [3]float64{u, vv, w})
	case "vn":
		// specifies a normal vector
		var i, j, k float64
		_, err = fmt.Sscanf(parts[1], "%f %f %f", &i, &j, &k)
		if err != nil {
			return
		}
		o.Vn = append(o.Vn, [3]float64{i, j, k})
	case "vp":
		var u, vv, w float64
		switch len(strings.Fields(parts[1])) {
		case 2:
			_, err = fmt.Sscanf(parts[1], "%f %f", &u, &vv)
		case 3:
			_, err = fmt.Sscanf(parts[1], "%f %f %f", &u, &vv, &w)
		default:
			err = fmt.Errorf("vp cannot handle %s", s)
		}
		if err != nil {
			return
		}
		o.Vp = append(o.Vp, [3]float64{u, vv, w})
		// Faces
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
				if s := strings.Split(coords, "/"); len(s[1]) == 0 {
					_, err = fmt.Sscanf(coords, "%d//%d", &a, &c)
				} else {
					_, err = fmt.Sscanf(coords, "%d/%d/%d", &a, &b, &c)
				}
				if err != nil {
					return
				}
				face = append(face, []int{a, b, c})
			}
		}
		o.F = append(o.F, face)
	}
	return nil
}

func Encode(w io.Writer, o Obj) error {
	// TODO
	return nil
}
