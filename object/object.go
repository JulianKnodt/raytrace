package object

import (
	v "github.com/julianknodt/vector"
)

/*
 Intersects: an object should have the
 ability to check whether it intersects with
 a vector

 Normal: an object should be able to determine a normal on its surface

 Material: an object should have a material for rendering
*/
type Object interface {
	// Surface Element maybe shouldn't be an interface, because the normal is constant...
	Intersects(v.Ray) (float64, SurfaceElement)
}

// Must implement this in order to even be considered a light source
type LightSource interface {
	// Whether it emits any light at all
	EmitsLight() bool
}
