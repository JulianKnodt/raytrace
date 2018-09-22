package scene

import (
	"os"

	v "github.com/julianknodt/vector"
	"raytrace/camera"
	"raytrace/color"
	"raytrace/lib/sky"
	"raytrace/obj"
	"raytrace/object"
	"raytrace/octree"
	"raytrace/off"
)

type Intersector func(v.Ray, Scene) *color.Normalized

type Scene struct {
	Height               float64 // This is resulting image height
	Width                float64 // This is resulting image width
	Objects              []object.Object
	Camera               camera.Camera
	Lights               []object.Object
	Sky                  sky.Sky
	IntersectionFunction Intersector
}

func (s *Scene) AddOff(filename string, shift v.Vec3) error {
	if filename == "" {
		return nil
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	model, err := off.Decode(f)
	if err != nil {
		return err
	}

	model.Vertices = v.Shift(model.Vertices, shift.X(), shift.Y(), shift.Z())
	s.Objects = append(s.Objects, octree.CreateFrom(model))
	return nil
}

func (s *Scene) AddObj(filename string, shift v.Vec3) error {
	if filename == "" {
		return nil
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	model, err := obj.Decode(f)
	if err != nil {
		return err
	}
	model.Shift(shift.X(), shift.Y(), shift.Z())
	s.Objects = append(s.Objects, octree.CreateFrom(model.IndexedTriangleList()))
	return nil
}

func (s *Scene) AddSky(filename string) error {
	if filename == "" {
		return nil
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	newSky, err := sky.NewSky(f)
	if err != nil {
		return err
	}

	s.Sky = *newSky
	return nil
}

func (s *Scene) AddLights() {
	if s.Lights == nil {
		s.Lights = []object.Object{}
	}
	for _, v := range s.Objects {
		if ls, ok := v.(object.LightSource); ok && ls.EmitsLight() {
			s.Lights = append(s.Lights, v)
		}
	}
}
