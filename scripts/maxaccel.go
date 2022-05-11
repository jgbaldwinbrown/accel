package main

import (
	"os"
	"../accel"
)

func main() {
	param := os.Args[1]
	accel.CalcAccelFull(os.Stdin, param)
}
