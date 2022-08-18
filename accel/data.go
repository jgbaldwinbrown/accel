package accel

import (
	"fmt"
	"io"
	"bufio"
	"github.com/jgbaldwinbrown/lscan/pkg"
	"strconv"
)

type Data struct {
	lines [][]string
	chrs []string
	chrxs map[string][]float64
	chrys map[string][]float64
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
