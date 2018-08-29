package mtl

import (
	"image"
	"path/filepath"
	"raytrace/utils"
	"sync"
)

type FileReference struct {
	FileName string
	Options  []string
	Args     []string
}

var loadedFiles sync.Map

func (f *FileReference) Load(mtlname string) (image.Image, error) {
	if f == nil {
		return nil, nil
	}

	file := filepath.Join(filepath.Dir(mtlname), f.FileName)
	if v, ok := loadedFiles.Load(file); ok {
		// always returns nil if memoized, presume nil
		return v.(image.Image), nil
	}
	img, err := utils.LoadImage(file)

	loadedFiles.Store(file, img)

	return img, err
}
