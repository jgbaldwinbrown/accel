package main

import (
	"strconv"
	"os"
	"github.com/jgbaldwinbrown/accel/accel"
)

func main() {
	param, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}
	accel.CalcQuantileFull(os.Stdin, param)
}
