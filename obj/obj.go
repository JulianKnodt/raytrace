package obj

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"raytrace/obj/mtl"
)

// Direct OBJ representation
type Obj struct {
	// Geometric Vertices
	V  [][3]float64
	Vt [][3]float64
	Vn [][3]float64
	Vp [][3]float64
	// Polygonal Face Element
	F []Face
	// Line Element
	L      [][]int
	MTLLib map[string]*mtl.MTL
}

type FaceElement struct {
	V  int
	Vt int
	Vn int
}

type Face struct {
	Elements []FaceElement
}

func Decode(file *os.File) (o *Obj, err error) {
	scanner := bufio.NewScanner(file)
	o = &Obj{
		V:      make([][3]float64, 1), // These are empty because they're 1 indexed
		Vt:     make([][3]float64, 1),
		Vn:     make([][3]float64, 1),
		F:      make([]Face, 0),
		MTLLib: make(map[string]*mtl.MTL),
	}

	currMTL := new(mtl.MTL)

	for scanner.Scan() {
		s := scanner.Text()

		parts := strings.SplitN(s, " ", 2)
		if len(parts) == 2 {
			parts[1] = strings.TrimSpace(parts[1])
		}
		switch parts[0] {
		case "", "#":
		case "g", "o":
		case "usemtl":
			if m, has := o.MTLLib[parts[1]]; has {
				*currMTL = *m
			}
		case "mtllib":
			mtlPath := filepath.Join(filepath.Dir(file.Name()), parts[1])
			mtlFile, err := os.Open(mtlPath)
			switch {
			case os.IsNotExist(err):
				continue
			case err != nil:
				return o, err
			}
			mtl, err := mtl.Decode(mtlFile)
			mtlFile.Close()
			for s, m := range mtl {
				o.MTLLib[s] = m
			}
		case "v":
			var v [3]float64
			v, err = parseFloats(parts[1])
			o.V = append(o.V, v)
		case "vt":
			var u, vv, w float64
			_, err = fmt.Sscanf(parts[1], "%f %f %f", &u, &vv, &w)
			o.Vt = append(o.Vt, [3]float64{u, vv, w})
		case "vn":
			// specifies a normal vector
			var i, j, k float64
			_, err = fmt.Sscanf(parts[1], "%f %f %f", &i, &j, &k)
			o.Vn = append(o.Vn, [3]float64{i, j, k})
		case "vp":
			var vp [3]float64
			vp, err = parseFloats(parts[1])
			o.Vp = append(o.Vp, vp)
		case "f":
			face := Face{}
			faceElements := make([]FaceElement, 0)
			for _, coords := range strings.Fields(parts[1]) {
				switch strings.Count(coords, "/") {
				case 0:
					var vertNumber int
					_, err = fmt.Sscanf(coords, "%d", &vertNumber)
					faceElements = append(faceElements, FaceElement{vertNumber, 0, 0})
				case 1:
					var a, b int
					_, err = fmt.Sscanf(coords, "%d/%d", &a, &b)
					faceElements = append(faceElements, FaceElement{a, b, 0})
				case 2:
					var a, b, c int
					if s := strings.Split(coords, "/"); len(s[1]) == 0 {
						_, err = fmt.Sscanf(coords, "%d//%d", &a, &c)
					} else {
						_, err = fmt.Sscanf(coords, "%d/%d/%d", &a, &b, &c)
					}
					faceElements = append(faceElements, FaceElement{a, b, c})
				}
			}
			face.Elements = faceElements
			o.F = append(o.F, face)
		}
		if err != nil {
			return
		}
	}

	return o, nil
}

func Encode(w io.Writer, o Obj) (err error) {
	// TODO
	return
}
