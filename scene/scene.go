package scene

import (
	"image/color"
	"os"

	"raytrace/camera"
	"raytrace/lib/sky"
	"raytrace/light"
	"raytrace/obj"
	"raytrace/object"
	"raytrace/off"
	v "raytrace/vector"
)

type Intersector func(v.Ray, Scene) color.Color

type Scene struct {
	Height               float64
	Width                float64
	Objects              []object.Object
	Camera               camera.Camera
	Lights               []light.Light
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
	s.Objects = append(s.Objects, model)
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
	s.Objects = append(s.Objects, model)
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
