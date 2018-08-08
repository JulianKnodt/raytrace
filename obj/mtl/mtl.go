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

// Representation of MTL file
type MTL struct {
	Ka [3]float64 // ambient reflectance // should be converted to its own type
	Kd [3]float64 // diffuse reflectance
	Ks [3]float64 // Specular reflectance
	D  float64    // Diffusion?
	Tf float64    // TODO transmission filter
	// Illumination mode of the MTL
	Illum Mode
	// name of the material
	Name string
}

func Decode(r io.Reader) (out map[string]MTL, err error) {
	out = make(map[string]MTL)
	scanner := bufio.NewScanner(r)
	var curr *MTL
	for scanner.Scan() {
		err, hasNew, newMTL := curr.addMTL(scanner.Text())
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
	case "sharpness":
	case "Ni", "ni":
	case "Material", "material":
	default:
		fmt.Println("Unmatched", s)
	}

	return
}
