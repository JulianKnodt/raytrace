package mtl

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	// ambient reflectance // should be converted to its own type
	Ka [3]float64
	// diffuse reflectance
	Kd [3]float64
	// Specular reflectance
	Ks [3]float64
	Ke [3]float64
	// Dissolve factor aka opacity
	D *float64
	// TODO transmission filter
	Tf float64
	// inverse of d
	Tr float64

	// Specular Exponent
	Ns float64
	// TODO label
	Ni float64
	// Illumination mode of the MTL
	Illum Mode
	// name of the material
	Name string

	// Mappings

	Map_Kd *FileReference
	/*
	   Specifies that a color texture file or a color procedural texture
	   is applied to the ambient reflectivity of the material.  During
	   rendering, the "map_Ka" value is multiplied by the "Ka" value.
	*/
	Map_Ka   *FileReference
	Map_Bump *FileReference
	fileName string
}

func CoerceMode(i Mode) Mode {
	switch {
	case i > 10:
		return 10
	case i < 0:
		return 0
	}
	return i
}

func Decode(file *os.File) (out map[string]*MTL, err error) {
	out = make(map[string]*MTL)
	scanner := bufio.NewScanner(file)
	m := new(MTL)
	m.fileName = file.Name()
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		switch parts := strings.SplitN(s, " ", 2); parts[0] {
		case "#", "": // Empty response
		case "newmtl":
			out[m.Name] = m
			m = new(MTL)
			m.fileName = file.Name()
		case "Ka", "ka":
			var x, y, z float64
			_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f %f %f", &x, &y, &z)
			m.Ka = [3]float64{x, y, z}
		case "Kd", "kd":
			var x, y, z float64
			_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f %f %f", &x, &y, &z)
			m.Kd = [3]float64{x, y, z}
		case "Ks", "ks":
			// https://stackoverflow.com/questions/36964747/ke-attribute-in-mtl-files
		case "Ke", "ke":
			var x, y, z float64
			_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f %f %f", &x, &y, &z)
			m.Ke = [3]float64{x, y, z}
		case "Tf", "tf":
		case "illum": // illumination
			var illum Mode
			_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &illum)
			m.Illum = CoerceMode(illum)
		case "d": // dissolve
			dissolve := new(float64)
			_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", dissolve)
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
			var transparency float64
			_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &transparency)
			m.Tr = transparency

		// Mapping

		case "map_Kd":
			// TODO
			switch fields := strings.Fields(parts[1]); len(fields) {
			case 1:
				// Just got the filename
				m.Map_Kd = &FileReference{
					FileName: fields[0],
				}
			default:
				// TODO Holy all the todos
			}

		case "map_Ka":
			switch fields := strings.Fields(parts[1]); len(fields) {
			case 1:
				m.Map_Ka = &FileReference{
					FileName: fields[0],
				}
			}
		case "map_bump", "bump":
			switch fields := strings.Fields(parts[1]); len(fields) {
			case 1:
				m.Map_Bump = &FileReference{
					FileName: fields[0],
				}
			}
		default:
			fmt.Println("Unmatched", s)
		}
		if err != nil {
			return out, err
		}
	}

	if m != nil {
		out[m.Name] = m
	}
	return
}

func (m MTL) Encode(w io.Writer) error {
	// TODO
	return nil
}
