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

const (
	ColorOnAmbientOff = iota
	ColorOnAmbientOn
	HighlightOn
	R_Raytrace
	T_Glass_R_Raytrace
	R_Fresnel_Raytrace
	T_Refraction_R_Raytrace
	T_Reflection_R_Fresnel_Raytrace
	T_Glass_R_RayTrace_Off
	Cast_Shadow // wtf is this
)

// Representation of MTL file
type MTL struct {
	// name of the material
	Name  string
	Ka    [3]float64
	Kd    [3]float64
	Ks    [3]float64
	Illum Mode
}

func Decode(r io.Reader) (out []MTL, err error) {
	out = make([]MTL, 0)
	curr := &MTL{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err, hasNew, newMTL := curr.addMTL(scanner.Text())
		if err != nil {
			return out, err
		}
		if hasNew {
			out = append(out, *curr)
			curr = newMTL
		}
	}
	return out, nil
}

func (m *MTL) addMTL(s string) (err error, hasNewMTL bool, newMTL *MTL) {
	switch parts := strings.SplitN(s, " ", 2); parts[0] {
	case "newmtl":
	case "Ka", "ka":
	case "Kd", "kd":
	case "Ks", "ks":
	case "Tf", "tf":
	case "illum": // illumination
	case "d": // dissolve
	case "Ns", "ns":
	case "sharpness":
	case "Ni", "ni":
	case "Material", "material":
	default:
		fmt.Println(s)
	}

	return
}
