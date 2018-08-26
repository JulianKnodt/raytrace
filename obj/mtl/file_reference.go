package mtl

import (
	"image"
	"path/filepath"
	"raytrace/utils"
)

type FileReference struct {
	FileName string
	Options  []string
	Args     []string
}

func (f *FileReference) Load(mtlname string) (image.Image, error) {
	if f == nil {
		return nil, nil
	}
	return utils.LoadImage(filepath.Join(filepath.Dir(mtlname), f.FileName))
}
