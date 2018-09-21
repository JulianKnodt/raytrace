package main

import (
	"math/rand"
	"sync"
	"time"

  "raytrace/scene"
  "raytrace/object"
  "raytrace/shapes"
  v "github.com/julianknodt/vector"
)

var rs = rand.New(rand.NewSource(time.Now().UnixNano()))

const numItems = 1000

func main() {
  objects := make([]object.Object, 0, numItems)
  for i := 0; i < numItems; i++ {
    objects = append(objects, shapes.NewSphere(v.RandomVector, rs.Float64(), nil))
  }
  scene := scene.Scene{
    Height: 1000.0,
    Width: 1200.0,
    IntersectionFunction: scene.Direct,
    Camera: camera.DefaultCamera(),
    Objects: objects,
  }


}
