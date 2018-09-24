package camera

import (
	v "github.com/julianknodt/vector"
)

var DefaultCameraDir = v.Vec3{0, 0, -1}

// Camera contains a transform as well as specific camera metadata
type Camera struct {
	// Location and Direction of Camera
	Transform *v.Ray

	// Represents the up direction of the camera
	Up v.Vec3

	// Angle pointing to the right of the camera
	// Computed and not meant to be changed directly
	Right v.Vec3

	// The distance that the camera renders pixels from the camera
	RenderDistance float64

	// The width and height of the rendered image
	Width, Height float64
}

func (c Camera) Location() v.Vec3 {
	return c.Transform.Origin
}

func (c Camera) Direction() v.Vec3 {
	return c.Transform.Direction
}

func NewCamera(position, direction, up v.Vec3,
	fieldOfView, renderDistance float64,
	width, height int,
) *Camera {
	return &Camera{
		Transform:      v.NewRay(position, direction),
		Up:             *up.Unit(),
		Right:          *direction.Cross(up).UnitSet(),
		RenderDistance: renderDistance,
		Width:          float64(width),
		Height:         float64(height),
	}
}

func (c Camera) AspectRatio() float64 {
	return c.Width / c.Height
}

func (c Camera) InverseWidth() float64 {
	return 1 / c.Width
}

func (c Camera) InverseHeight() float64 {
	return 1 / c.Height
}

func DefaultCamera() Camera {
	return Camera{
		Transform:      v.NewRay(v.Vec3{0, 0, 0}, v.Vec3{0, 0, -1}),
		Up:             v.Vec3{0, 1, 0},
		Right:          v.Vec3{-1, 0, 0},
		RenderDistance: 1,
		Width:          5, // This is world space
		Height:         5, // This is world space
	}
}

// Returns ray from camera position to [0,1], [0, 1] in its viewing box
func (c Camera) RayTo(x, y float64) v.Ray {

	// have to divide by 2 since it extends in both directions
	hComp := c.Right.SMul((1 - 2*x) * c.Width / 2)
	vComp := c.Up.SMul((1 - 2*y) * c.Height / 2)

	return *v.NewRay(c.Transform.Origin,
		*c.Transform.Direction.SMul(c.RenderDistance).
			AddSet(*hComp).
			AddSet(*vComp),
	)
}

// Convenience for seeing what the box that the camera can see is
func (c Camera) Range() (min, max v.Vec3) {
	min = *c.Transform.Origin.
		Add(*c.Transform.Direction.SMul(c.RenderDistance)).
		AddSet(*c.Right.Inv().SMulSet(c.Width / 2)).
		AddSet(*c.Up.Inv().SMulSet(c.Height / 2))
	max = *c.Transform.Origin.
		Add(*c.Transform.Direction.SMul(c.RenderDistance)).
		AddSet(*c.Right.SMul(c.Width / 2)).
		AddSet(*c.Up.SMul(c.Height / 2))
	return
}
