package main

import (
	"os"
	"github.com/jgbaldwinbrown/accel/accel"
)

func main() {
	param := os.Args[1]
	accel.CalcAccelFull(os.Stdin, param)
}
