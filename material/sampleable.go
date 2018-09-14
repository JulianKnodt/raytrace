package material

import (
	v "github.com/julianknodt/vector"
)

// Sampleable is meant to be implemented by surface elements
// Which return a uv coordinate for mapping onto a texture map
type Sampleable interface {

	// Returns sampleable coordinates for the given vector
	// Should both be in [0,1]
	TextureCoordinates(v.Vec3) (u, v float64)
}
