package octree

import (
  v "raytrace/vector"
)

type Octree struct {
  Center v.Vec3
  // Far X Far Y Far Z
  XYZ v.Vec3
  XYNZ v.Vec3
  XNYZ v.Vec3
  XNYNZ v.Vec3
  NXYZ v.Vec3
  NXYNZ v.Vec3
  NXNYZ v.Vec3
  NXNYNZ v.Vec3
}
