package obj

import (
	"strconv"
	"strings"
)

func parseFloats(s string) ([3]float64, error) {
	out := make([]float64, 4) // can possibly have 4 elements
	for i, v := range strings.Fields(s) {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return [3]float64{}, err
		}
		out[i] = f
	}
	return [3]float64{out[0], out[1], out[2]}, nil
}
