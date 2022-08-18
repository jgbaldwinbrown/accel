package accel

import (
	"fmt"
	"github.com/chewxy/stl/loess"
	"math"
	"io"
)

func CalcAccel(data []float64, distf func(p float64) float64) (maxidx int, maxy, maxaccel float64, err error) {
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
		err = fmt.Errorf("Empty file.")
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
	return maxi, data[maxi], max, nil
}

func AccelDistf(param string) func(float64) float64 {
	var distf func(float64)float64
	switch param {
	case "fst":
		distf = ExpDistf()
	case "selec":
		distf = ExpDistf()
	case "pfst":
		distf = ExpDistf()
	default:
		distf = func(p float64) float64 { return p }
	}
	return distf
}

func AccelCol(param string) int {
	if param == "pfst" {
		return 8
	}
	return 3
}

func CalcAccelFull(r io.Reader, param string) {
	col := AccelCol(param)
	distf := AccelDistf(param)

	data, err := ReadTable(r, col)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxi, maxy, maxaccel, err := CalcAccel(data, distf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(maxi, maxy, maxaccel)
}
