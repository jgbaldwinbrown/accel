package main

import (
	"errors"
	"os"
	"io"
	"fmt"
	"github.com/jgbaldwinbrown/lscan/lscan"
	"bufio"
	"strconv"
)

type Data struct {
	lines [][]string
	chrs []string
	chrxs map[string][]float64
	chrys map[string][]float64
}

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

func ReadLines(r io.Reader) (Data, error) {
	var data Data
	data.chrxs = make(map[string][]float64)
	data.chrys = make(map[string][]float64)
	s := bufio.NewScanner(r)
	s.Buffer([]byte{}, 1e12)
	splitter := lscan.ByByte('\t')
	for s.Scan() {
		var line []string
		line = lscan.SplitByFunc(line, s.Text(), splitter)
		chr := line[0]

		data.lines = append(data.lines, line)

		left, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return data, err
		}
		right, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return data, err
		}
		x := (left + right-1) / 2
		y, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			return data, err
		}

		_, ok := data.chrxs[chr]
		if !ok {
			data.chrs = append(data.chrs, chr)
			data.chrxs[chr] = []float64{}
		}
		data.chrxs[chr] = append(data.chrxs[chr], x)

		_, ok = data.chrys[chr]
		if !ok {
			data.chrys[chr] = []float64{}
		}
		data.chrys[chr] = append(data.chrys[chr], y)
	}
	return data, nil
}

func DerivAll(d Data) (Data, error) {
	var out Data
	out.lines = d.lines
	out.chrs = d.chrs
	out.chrxs = make(map[string][]float64)
	out.chrys = make(map[string][]float64)
	for _, chr := range d.chrs {
		out.chrxs[chr], out.chrys[chr], _ = Deriv(d.chrxs[chr], d.chrys[chr])
	}
	return out, nil
}

func WriteAll(w io.Writer, d Data) {
	for _, chr := range d.chrs {
		xs := d.chrxs[chr]
		ys := d.chrys[chr]
		for i, x := range xs {
			y := ys[i]
			fmt.Fprintf(w, "%s\t%d\t%d\t%g\n", chr, int(x), int(x+1), y)
		}
	}
}

func main() {
	data, err := ReadLines(os.Stdin)
	if err != nil {
		panic(err)
	}

	d1, err := DerivAll(data)
	if err != nil {
		panic(err)
	}
	d2, err := DerivAll(d1)
	if err != nil {
		panic(err)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	WriteAll(out, d2)
}

// chr1    0       50000   0.00803825758730159     1       1000
// chr1    5000    55000   0.007246008738461539    1       6000
// chr1    10000   60000   0.00618197719047619     1       11000
// chr1    15000   65000   0.0063951448288288295   1       16000
// chr1    20000   70000   0.008938101742857142    1       21000
// chr1    25000   75000   0.0048611521509433955   1       26000
// chr1    30000   80000   0.004164311958333331    1       31000
// chr1    35000   85000   0.005828132695652173    1       36000
// chr1    40000   90000   0.006530389738095237    1       41000
// chr1    45000   95000   0.003951702302325582    1       46000
// chr1    50000   100000  0.0007976492705882357   1       51000
// chr1    55000   105000  0.0010451914880952374   1       56000
// chr1    60000   110000  0.001225674101123596    1       61000
// chr1    65000   115000  0.0024636257926829263   1       66000
// chr1    70000   120000  0.0014977833132530118   1       71000
// chr1    75000   125000  0.005601852901234561    1       76000
// chr1    80000   130000  0.016465384249999996    1       81000
// chr1    85000   135000  0.021192385764705876    1       86000
// chr1    90000   140000  0.023885360804597692    1       91000
// chr1    95000   145000  0.021090373695652167    1       96000
// chr1    100000  150000  0.02171647838383838     1       101000
// chr1    105000  155000  0.026366671960784308    1       106000
