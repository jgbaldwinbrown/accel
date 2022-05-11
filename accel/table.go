package accel

import (
	"github.com/jgbaldwinbrown/lscan/lscan"
	"bufio"
	"sort"
	"io"
	"strconv"
	"errors"
)

func ReadTable(r io.Reader, col int) ([]float64, error) {
	var data []float64
	var line []string
	s := bufio.NewScanner(r)
	s.Buffer([]byte{}, 1e12)
	splitter := lscan.ByByte('\t')
	for s.Scan() {
		line = lscan.SplitByFunc(line, s.Text(), splitter)
		f, err := strconv.ParseFloat(line[col], 64)
		if err != nil {
			return data, errors.New("Line too short.")
		}
		data = append(data, f)
	}
	sort.Float64s(data)
	return data, nil
}
