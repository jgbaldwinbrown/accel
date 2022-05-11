package accel

import (
	"errors"
)

func Deriv(x []float64, y []float64) (xout, yout []float64, err error) {
	if len(x) < 1 {
		return xout, yout, errors.New("Deriv error: too short.")
	}
	if len(x) != len(y) {
		return xout, yout, errors.New("Deriv error: lengths do not match.")
	}
	xout = make([]float64, len(x) - 1)
	yout = make([]float64, len(y) - 1)
	l := len(x) - 1
	for i:=0; i<l; i++ {
		xout[i] = (x[i] + x[i+1]) / 2
		yout[i] = (y[i+1] - y[i]) / (x[i+1] - x[i])
	}
	return xout, yout, nil
}
