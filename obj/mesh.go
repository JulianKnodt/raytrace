package obj

import (
	"raytrace/material"
	"raytrace/mesh"
)

func (o Obj) Mesh() *mesh.Mesh {
	materials := make(map[string]*material.Material, len(o.MTLLib))
	for k, v := range o.MTLLib {
		materials[k] = v.Material()
	}

	return nil
}
