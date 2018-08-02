package off

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"raytrace/mesh"
	v "raytrace/vector"
	"strconv"
	"strings"
)

func getTriple(s string) (out v.Vec3, err error) {
	parts := strings.Split(s, " ")
	n, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return out, err
	}
	out[0] = n
	n, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return out, err
	}
	out[1] = n
	n, err = strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return out, err
	}
	out[2] = n
	return out, nil
}

func Decode(offFile *os.File) (*mesh.Mesh, error) {
	scanner := bufio.NewScanner(offFile)
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
		vec, err := getTriple(scanner.Text())
		if err != nil {
			return nil, err
		}
		result.Vertices = append(result.Vertices, vec)
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

func Encode(offFile *os.File, m mesh.Mesh) error {
	if err := offFile.Truncate(0); err != nil {
		return err
	}

	if _, err := offFile.Write([]byte("OFF\n")); err != nil {
		return err
	}

	if _, err := offFile.Write([]byte(fmt.Sprintf("%d %d %d\n", m.Verts(), m.Faces(), m.Edges()))); err != nil {
		return err
	}

	for _, v := range m.Vertices {
		if _, err := offFile.Write([]byte(fmt.Sprintf("%f %f %f\n", v[0], v[1], v[2]))); err != nil {
			return err
		}
	}

	for _, order := range m.Order {
		if _, err := offFile.Write([]byte(fmt.Sprintf("%d  %s", len(order), strings.Join(intArrToStringArr(order), " ")))); err != nil {
			return err
		}
	}

	return nil
}
