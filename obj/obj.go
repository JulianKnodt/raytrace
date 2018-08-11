package obj

import (
	"bufio"
	"fmt"
	"io"
	"raytrace/obj/mtl"
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
	F []Face
	// Line Element
	L [][]int
}

type FaceElement struct {
	V  int
	Vt int
	Vn int
}

type Face struct {
	Elements []FaceElement
}

type State struct {
	G         string
	O         string
	UseMTL    string
	MTLLib    map[string]mtl.MTL
	MTLLoader MTLLoader
}

type MTLLoader func(string) (map[string]mtl.MTL, error)

func Decode(r io.Reader, mtlLoader MTLLoader) (*Obj, error) {
	scanner := bufio.NewScanner(r)
	o := &Obj{
		V:  make([][4]float64, 1), // These are empty because they're 1 indexed
		Vt: make([][3]float64, 1),
		Vn: make([][3]float64, 1),
		F:  make([]Face, 0),
	}

	state := &State{
		MTLLib:    make(map[string]mtl.MTL),
		MTLLoader: mtlLoader,
	}

	for scanner.Scan() {
		if err := o.add(scanner.Text(), state); err != nil {
			return o, err
		}
	}
	return o, nil
}

func (o *Obj) add(s string, state *State) (err error) {
	parts := strings.SplitN(s, " ", 2)
	if len(parts) == 2 {
		parts[1] = strings.TrimSpace(parts[1])
	}
	switch parts[0] {
	case "", "#":
	case "g":
		state.G = strings.TrimSpace(parts[1])
	case "o":
	case "usemtl":
		state.UseMTL = strings.TrimSpace(parts[1])
	case "mtllib":
		if state.MTLLoader == nil {
			// skipping loading mtl files
			return nil
		}
		mtl, err := state.MTLLoader(strings.TrimSpace(parts[1]))
		if err != nil {
			return err
		}
		for s, m := range mtl {
			state.MTLLib[s] = m
		}
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
		face := Face{}
		faceElements := make([]FaceElement, 0)
		for _, coords := range strings.Fields(parts[1]) {
			switch strings.Count(coords, "/") {
			case 0:
				var vertNumber int
				_, err := fmt.Sscanf(coords, "%d", &vertNumber)
				if err != nil {
					return err
				}
				faceElements = append(faceElements, FaceElement{vertNumber, -1, -1})
			case 1:
				var a, b int
				_, err := fmt.Sscanf(coords, "%d/%d", &a, &b)
				if err != nil {
					return err
				}
				faceElements = append(faceElements, FaceElement{a, b, -1})
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
				faceElements = append(faceElements, FaceElement{a, b, c})
			}
		}
		face.Elements = faceElements
		o.F = append(o.F, face)
	}
	return nil
}

func Encode(w io.Writer, o Obj) (err error) {
	// TODO
	return
}
