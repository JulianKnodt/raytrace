package off

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"raytrace/mesh"
	v "raytrace/vector"
	"strconv"
	"strings"
)

func Decode(r io.Reader) (*mesh.Mesh, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if scanner.Text() != "OFF" {
		return nil, errors.New("Not an OFF file")
	}

	var numVertices uint64
	var numFaces uint64

	result := new(mesh.Mesh)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix("#", strings.TrimLeft(text, " ")) {
			continue
		} else {
			parts := strings.Split(text, " ")
			n, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}
			numVertices = uint64(n)
			n, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
			numFaces = uint64(n)
			break
		}
	}

	for i := uint64(0); i < numVertices; i++ {
		scanner.Scan()
		var x, y, z float64
		_, err := fmt.Sscanf(scanner.Text(), "%f %f %f", &x, &y, &z)
		if err != nil {
			return nil, err
		}
		result.Vertices = append(result.Vertices, v.Vec3{x, y, z})
	}

	for i := uint64(0); i < numFaces; i++ {
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		numVert, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		order := make([]int, numVert)
		for _, v := range parts[1:] {
			if len(v) > 0 {
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				order = append(order, n)
			}
		}
		result.Order = append(result.Order, order)
	}

	return result, nil
}

func intArrToStringArr(a []int) []string {
	result := make([]string, 0, len(a))
	for _, v := range a {
		result = append(result, string(v))
	}
	return result
}

func Encode(w io.Writer, m mesh.Mesh) error {
	if _, err := fmt.Fprint(w, "OFF\n"); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "%d %d %d\n", m.Verts(), m.Faces(), m.Edges()); err != nil {
		return err
	}

	for _, v := range m.Vertices {
		if _, err := fmt.Fprintf(w, "%f %f %f\n", v[0], v[1], v[2]); err != nil {
			return err
		}
	}

	for _, order := range m.Order {
		if _, err := fmt.Fprintf(w, "%d  %s\n", len(order), strings.Join(intArrToStringArr(order), " ")); err != nil {
			return err
		}
	}

	return nil
}
