package obj

// Shifts an obj model by x, y, z, with weight w
func (o *Obj) Shift(x, y, z, w float64) {
	for i, _ := range o.V {
		o.V[i][0] += x
		o.V[i][1] += y
		o.V[i][2] += z
		o.V[i][3] += w
	}
}
