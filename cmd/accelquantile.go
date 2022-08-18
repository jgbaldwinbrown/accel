package main

import (
	"strconv"
	"os"
	"github.com/jgbaldwinbrown/accel/accel"
)

func main() {
	accel_param := os.Args[1]
	quantile, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		panic(err)
	}

	err = accel.CalcQuantileAndAccelFull(os.Stdin, accel_param, quantile)
	if err != nil {
		panic(err)
	}
}
