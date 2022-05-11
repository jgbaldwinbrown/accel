package accel

import (
	"fmt"
	"github.com/chewxy/stl/loess"
	"math"
	"io"
)

func CalcAccelFull(r io.Reader, param string) {
	var distf func(float64)float64
	col := 3
	switch param {
	case "fst":
		distf = ExpDistf()
	case "selec":
		distf = ExpDistf()
	case "pfst":
		col = 8
		distf = ExpDistf()
	default:
		distf = func(p float64) float64 { return p }
	}

	data, err := ReadTable(r, col)
	if err != nil {
		fmt.Println(err)
		return
	}

	y, err := loess.Smooth(data, 3000, 1, loess.Linear)
	if err != nil {
		panic(err)
	}

	x := DistribIndices(len(data), distf)

	lx := len(x)
	ly := len(y)
	x = x[lx/4:lx-(lx/20000)]
	y = y[ly/4:ly-(ly/20000)]

	xd1, yd1, _ := Deriv(y, x)
	_, yd2, _ := Deriv(xd1, yd1)

	if len(yd2) < 1 {
		fmt.Println("Empty file.")
		return
	}
	max := yd2[0]
	maxi := 0
	for i, y := range yd2[1:] {
		if math.IsNaN(max) || math.IsInf(max, 0) || (max < y && !math.IsInf(y, 0)) {
			max = y
			maxi = i
		}
	}
	fmt.Println(data[maxi], max)
}
