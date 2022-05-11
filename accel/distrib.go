package accel

import (
	"gonum.org/v1/gonum/stat/distuv"
)

func DistribIndices(n int, f func(float64) float64) []float64 {
	out := make([]float64, n)
	fn := float64(n)

	for i:=0; i<n; i++ {
		out[i] = f(float64(i)/fn)
	}
	return out
}

func ExpDistf() func(p float64) float64 {
	dist := distuv.Exponential{}
	dist.Rate = 12.5
	return func(p float64) float64 {
		return dist.Quantile(p)
	}
}
