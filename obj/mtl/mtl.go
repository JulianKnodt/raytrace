package mtl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Represents the illumination mode of
// this file
type Mode int8

func (m Mode) String() string {
	switch m {
	case 0:
		return "ColorOnAmbientOff"
	case 1:
		return "ColorOnAmbientOn"
	case 2:
		return "HighlightOn"
	case 3:
		return "R_Raytrace"
	case 4:
		return "T_Glass_R_Raytrace"
	case 5:
		return "R_Fresnel_Raytrace"
	case 6:
		return "T_Refraction_R_Raytrace"
	case 7:
		return "T_Reflection_R_Fresnel_Raytrace"
	case 8:
		return "Reflection_RaytraceOff"
	case 9:
		return "T_Glass_R_RayTrace_Off"
	case 10:
		return "Cast_Shadow"
	}
	return "Unknown"
}

const (
	ColorOnAmbientOff = iota
	ColorOnAmbientOn
	HighlightOn
	R_Raytrace
	T_Glass_R_Raytrace
	R_Fresnel_Raytrace
	T_Refraction_R_Raytrace
	T_Reflection_R_Fresnel_Raytrace
	Reflection_RaytraceOff
	T_Glass_R_RayTrace_Off
	Cast_Shadow // wtf is this
)

type Map_Kd struct {
	FileName string
	Options  []string
	Args     []string
}

// Representation of MTL file
type MTL struct {
	// ambient reflectance // should be converted to its own type
	Ka [3]float64
	// diffuse reflectance
	Kd [3]float64
	// Specular reflectance
	Ks [3]float64
	Ke [3]float64
	// Diffusion?
	D float64
	// TODO transmission filter
	Tf float64
	// TODO label // is this even in the spec???
	Tr [3]float64
	// TODO label
	Ns float64
	// TODO label
	Ni float64
	// Illumination mode of the MTL
	Illum Mode
	// name of the material
	Name string

	// Mappings

	Map_Kd Map_Kd
}

func Decode(r io.Reader) (out map[string]MTL, err error) {
	out = make(map[string]MTL)
	scanner := bufio.NewScanner(r)
	var curr *MTL
	for scanner.Scan() {
		err, hasNew, newMTL := curr.addMTL(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return out, err
		} else if hasNew {
			if curr != nil {
				out[curr.Name] = *curr
			}
			curr = newMTL
		}
	}

	if curr != nil {
		out[curr.Name] = *curr
	}
	return
}

func (m *MTL) addMTL(s string) (err error, hasNewMTL bool, newMTL *MTL) {
	switch parts := strings.SplitN(s, " ", 2); parts[0] {
	case "#", "": // Empty response
	case "newmtl":
		hasNewMTL = true
		newMTL = &MTL{
			Name: strings.TrimSpace(parts[1]),
		}
	case "Ka", "ka":
		var x, y, z float64
		_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f %f %f", &x, &y, &z)
		m.Ka = [3]float64{x, y, z}
	case "Kd", "kd":
		var x, y, z float64
		_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f %f %f", &x, &y, &z)
		m.Kd = [3]float64{x, y, z}
	case "Ks", "ks":
	case "Ke", "ke":
		var x, y, z float64
		_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f %f %f", &x, &y, &z)
		m.Ke = [3]float64{x, y, z}
	case "Tf", "tf":
	case "illum": // illumination
		var illum Mode
		_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &illum)
		m.Illum = illum
	case "d": // dissolve
		var dissolve float64
		_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &dissolve)
		m.D = dissolve
	case "Ns", "ns":
		var ns float64
		_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &ns)
		m.Ns = ns
	case "sharpness":
	case "Ni", "ni":
		var ni float64
		_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &ni)
		m.Ni = ni
	case "Material", "material":
	case "Tr", "tr":
		// TODO

	// Mapping

	case "map_Kd":
		// TODO
		fields := strings.Fields(parts[1])
		switch len(fields) {
		case 1:
			// Just got the filename
			m.Map_Kd = Map_Kd{
				FileName: fields[0],
			}
		default:
			// TODO Holy all the todos
		}

	default:
		fmt.Println("Unmatched", s)
	}

	return
}
