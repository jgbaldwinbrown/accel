package accel

import (
	"reflect"
	"testing"
)

type qTest struct {
	data []float64
	quantile float64
	dir int
	threshpos int
	thresh float64
	overthresh []float64
}

func TestQuantile(t *testing.T) {
	var data, data2 []float64
	for i:=0; i<100; i++ {
		data = append(data, float64(i))
	}
	for i:=99; i>=0; i-- {
		data2 = append(data2, float64(i))
	}
	tests := []qTest{
		qTest{data, .5, 1, 50, 50.0, data[50:]},
		qTest{data, .1, 1, 90, 90.0, data[90:]},
		qTest{data2, .5, 1, 50, 50.0, data[50:]},
		qTest{data2, .1, 1, 90, 90.0, data[90:]},
		qTest{data, .5, -1, 50, 49.0, data2[50:]},
		qTest{data, .1, -1, 90, 9.0, data2[90:]},
		qTest{data2, .5, -1, 50, 49.0, data2[50:]},
		qTest{data2, .1, -1, 90, 9.0, data2[90:]},
	}
	for _, test := range tests {
		tout := test
		tout.threshpos, tout.thresh, tout.overthresh = internalQuantile(test.data, test.quantile, test.dir)
		if !reflect.DeepEqual(test, tout) {
			t.Errorf("test %v and tout %v differ", test, tout)
		}
	}
}
