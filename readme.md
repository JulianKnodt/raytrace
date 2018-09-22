# raytrace

A raytracer implemented in golang for rendering off and obj files.

It uses octrees and bounding boxes in order to optimize the speed with which intersections
must be checked, and also allows for more implementations of intersection algorithms.

It was implemented using the [Graphics Codex](http://graphicscodex.com/index.php) as a reference
as well as various other sources.

The list of sources can be found by running `make only-sources`

And tests can be run by running `make test` and benches are similarly `make bench`

If you want to contribute, todos can be found by `make todos`

---

# Sample Outputs

![Dragon](./outputs/dragon.png)
![Teapot](./outputs/teapot.png)




---

## Sources

- https://computergraphics.stackexchange.com/questions/375/what-is-ambient-lighting
- https://learnopengl.com/Getting-started/Textures https://www.cs.cmu.edu/afs/cs/academic/class/15462-f09/www/lec/lec8.pdf
- https://cglearn.codelight.eu/pub/computer-graphics/textures-and-sampling
- https://stackoverflow.com/questions/36964747/ke-attribute-in-mtl-files
- http://ogldev.atspace.co.uk/www/tutorial16/tutorial16.html
- Implementation of https://en.wikipedia.org/wiki/Wavefront_.obj_file#Line_elements
- http://realtimecollisiondetection.net/books/rtcd/
- https://github.com/marczych/RayTracer
- http://graphics.cs.williams.edu/data
- https://gamedev.stackexchange.com/questions/114412/how-to-get-uv-coordinates-for-sphere-cylindrical-projection
- https://github.com/fogleman/pt/blob/master/pt/sphere.go
- https://www.quora.com/What-is-the-difference-between-Ambient-Diffuse-and-Specular-Light-in-OpenGL-Figures-for-illustration-are-encouraged
- http://cse.csusb.edu/tongyu/courses/cs520/notes/texture.php
- https://people.cs.clemson.edu/~dhouse/courses/405/docs/brief-mtl-file-format.html
- http://blog.lexique-du-net.com/index.php?post/2009/07/24/AmbientDiffuseEmissive-and-specular-colorSome-examples
- http://paulbourke.net/dataformats/mtl/
- https://gamedev.stackexchange.com/questions/23743/whats-the-most-efficient-way-to-find-barycentric-coordinates
- https://en.wikipedia.org/wiki/Shading#Ambient_lighting
