package bounding

import (
	"fmt"
)

func (a AxisAlignedBoundingBox) String() string {
	return fmt.Sprintf("AABB { \n\tx: %f - %f,\n\ty: %f - %f,\n\tz: %f - %f}",
		a.Xx, a.XX, a.Yy, a.YY, a.Zz, a.ZZ)
}
