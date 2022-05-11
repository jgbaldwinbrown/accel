package main

import (
	"os"
	"bufio"
	"../accel"
)

func main() {
	data, err := accel.ReadLines(os.Stdin)
	if err != nil {
		panic(err)
	}

	d1, err := accel.DerivAll(data)
	if err != nil {
		panic(err)
	}
	d2, err := accel.DerivAll(d1)
	if err != nil {
		panic(err)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	accel.WriteAll(out, d2)
}
