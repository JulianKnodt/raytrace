package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	// "image/gif"
	"image/png"

	v "github.com/julianknodt/vector"
	"raytrace/bounding"
	"raytrace/camera"
	"raytrace/color"
	"raytrace/material"
	"raytrace/object"
	"raytrace/octree"
	"raytrace/scene"
	"raytrace/shapes"
	"raytrace/stl"
	"time"
)

var (
	out  = flag.String("o", "out.png", "The location of the resulting gif")
	comp = flag.String("comp", "", "The location of the direct rendering of the stl file")

	stlFile = flag.String("stl", "", "The stl file to be be used as a reference")
	isAscii = flag.Bool("ascii", false, "Use Ascii STL file parsing?")

	thres = flag.Float64("t", 0.01, `Max area threshold allowed for triangles relative to the
  distance between the corners of the bounding box for the whole shape. Should be between [0,1]
  but can be bigger than 1.`)
	vecPer  = flag.Float64("yield", 0.05, "Percent of vectors to use as triangles while rendering")
	scaling = flag.Float64("scale", 1, "How much to scale the image so that it fits within a 1x1 box")

	seed     = flag.Int64("seed", time.Now().UnixNano(), "Seed to use when rendering the image")
	showSeed = flag.Bool("show-seed", false, "Whether or not to show random seed")

	shiftX = flag.Float64("sx", 0, "Amt to shift in the x direction")
	shiftY = flag.Float64("sy", 0, "Amt to shift in the y direction")
	shiftZ = flag.Float64("sz", 0, "Amt to shift in the z direction")

	rotateX = flag.Float64("rx", 0, "Amt to rotate in the x direction")
	rotateY = flag.Float64("ry", 0, "Amt to rotate in the y direction")
	rotateZ = flag.Float64("rz", 0, "Amt to rotate in the z direction")
)

func ValidTriple(a, b, c v.Vec3, maxArea float64) bool {
	return v.TriangleArea(a, b, c) < maxArea
	/*		a.Sub(b).SqrMagn() < 0.75 && a.Sub(b).SqrMagn() > 0.005 &&
			c.Sub(b).SqrMagn() < 0.75 && c.Sub(b).SqrMagn() > 0.005 &&
			a.Sub(c).SqrMagn() < 0.75 && a.Sub(c).SqrMagn() > 0.005
	*/
}

func clamp(a float64) float64 {
	switch {
	case a < 0:
		return 0
	case a > 1:
		return 1
	}
	return a
}

func AddOp(b float64) func(float64) float64 {
	return func(a float64) float64 {
		return a + b
	}
}

func main() {
	flag.Parse()
	switch {
	case *out == "":
		fmt.Println("Cannot pass empty output file")
		return
	case *stlFile == "":
		fmt.Println("Must reference stl file with -stl")
		return
	case *thres < 0:
		fmt.Println("Must have positive threshold")
		return
	case *scaling == 0:
		fmt.Println("Scaling to 0 is kinda pointless.")
		return
	}

	if *showSeed {
		fmt.Println("Seed: ", *seed)
	}
	var rs = rand.New(rand.NewSource(*seed))

	// Take only vertex data from stlFile
	f, err := os.Open(*stlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var original []stl.Triangle
	if *isAscii {
		original, err = stl.DecodeAscii(f)
	} else {
		original, err = stl.DecodeBinary(f)
	}
	fmt.Printf("Original has %d triangles\n", len(original))

	if err != nil {
		panic(err)
	}
	vectors := make([]v.Vec3, 0, len(original)*3)
	for _, vec := range original {
		v1, v2, v3 := *vec.V1.SMul(*scaling), *vec.V2.SMul(*scaling), *vec.V3.SMul(*scaling)
		vectors = append(vectors, v1, v2, v3)
	}

	//vectors = v.Rotate(vectors, *rotateX, *rotateY, *rotateZ)
	vectors = v.Shift(vectors, *shiftX, *shiftY, *shiftZ)

	min, max := v.Inf(1), v.Inf(-1)
	for _, vec := range vectors {
		min.MinSet(vec)
		max.MaxSet(vec)
	}

	c := camera.DefaultCamera()
	c.Width = 2
	c.Height = 2
	c.Transform.Origin[2] = 2

	if *comp != "" {
		box := bounding.Box{Min: *min, Max: *max}
		oct := octree.NewEmptyOctree(box)
		ts := stl.ToTriangles(original, &material.Material{
			Ambient: color.DefaultColor,
		})
		items := make([]octree.OctreeItem, len(ts))
		for i, t := range ts {
			items[i] = t
		}
		oct.Insert(items...)
		scene := scene.Scene{
			Height:               600.0,
			Width:                600.0,
			IntersectionFunction: scene.Direct,
			Camera:               c,
			Objects:              []object.Object{oct},
			Lights: []object.LightSource{
				shapes.NewSphere(v.Vec3{-12, 0, 12}, 1, material.WhiteLightMaterial()),
			},
		}
		img := scene.Render()
		f, err := os.Create(*comp)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err = png.Encode(f, img); err != nil {
			panic(err)
		}
	}

	fmt.Println("STL within ", *min, *max)
	maxArea := min.Sub(*max).Magn() * (*thres)
	fmt.Println("Max area is: ", maxArea)

	// Randomly add triangles between vertices between certain threshold of area and/or length
	// of various colours

	numTriangles := len(vectors)
	maxVecs := int(clamp(*vecPer) * float64(numTriangles))
	fmt.Printf("Rendering %d triangles\n", maxVecs)
	triangles := make([]octree.OctreeItem, 0, maxVecs)

	for i := 0; i < maxVecs; i++ {
		for {
			a := i
			for a == i {
				a = rs.Intn(numTriangles)
			}
			b := i
			for b == i || a == b {
				b = rs.Intn(numTriangles)
			}
			if ValidTriple(vectors[a], vectors[b], vectors[i], maxArea) {
				mat := &material.Material{
					Ambient:    colors[rs.Intn(len(colors))],
					RenderType: material.Lambertian{Albedo: 1},
				}
				triangles = append(triangles, shapes.NewTriangle(vectors[a], vectors[b], vectors[i], mat))
				break
			}
		}
	}
	box := bounding.Box{Min: *min, Max: *max}
	oct := octree.NewEmptyOctree(box)
	oct.Insert(triangles...)

	fmt.Println("Finished finding triangles: Camera looking at")
	fmt.Println(c.Range())
	scene := scene.Scene{
		Height:               600.0,
		Width:                600.0,
		IntersectionFunction: scene.Direct,
		Camera:               c,
		Objects:              []object.Object{oct},
		Lights: []object.LightSource{
			shapes.NewSphere(v.Vec3{-12, 0, 12}, 1, material.WhiteLightMaterial()),
		},
	}
	// render image

	img := scene.Render()
	output, err := os.Create(*out)
	if err != nil {
		panic(err)
	}
	defer output.Close()
	err = png.Encode(output, img)
	if err != nil {
		panic(err)
	}

	// Repeat a couple of times to produce a gif
	// TODO
}
