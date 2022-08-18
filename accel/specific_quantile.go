package accel

import (
	"fmt"
	"math"
	"io"
)

func Quantile(data []float64, quantile float64) (threshpos int, thresh float64, overthresh []float64) {
	ly := float64(len(data))
	threshpos = int(math.Ceil(ly * (1 - quantile)))
	thresh = data[threshpos]
	overthresh = data[threshpos:]
	return
}

func CalcQuantileFull(r io.Reader, param float64) {
	col := 3
	y, err := ReadTable(r, col)
	if err != nil {
		fmt.Println(err)
		return
	}

	threshpos, thresh, overthresh := Quantile(y, param)

	if len(overthresh) < 1 {
		fmt.Println("Empty file.")
		return
	}
	fmt.Println(param, threshpos, thresh, len(overthresh))
}

type Stats struct {
	MaxIdx int
	MaxY float64
	MaxAccel float64
	ThreshPos int
	Thresh float64
	OverThresh []float64
}

const StatsHeader string = "MaxIdx\tMaxY\tMaxAccel\tThreshPos\tThresh\tNoverThresh"

func (s Stats) Pretty() string {
	return fmt.Sprintf(
		"%v\t%v\t%v\t%v\t%v\t%v",
		s.MaxIdx,
		s.MaxY,
		s.MaxAccel,
		s.ThreshPos,
		s.Thresh,
		len(s.OverThresh),
	)
}

func CalcQuantileAndAccelFull(r io.Reader, accelparam string, quantilethresh float64) error {
	col := AccelCol(accelparam)
	distf := AccelDistf(accelparam)
	y, err := ReadTable(r, col)
	if err != nil {
		return err
	}

	var s Stats
	s.ThreshPos, s.Thresh, s.OverThresh = Quantile(y, quantilethresh)

	s.MaxIdx, s.MaxY, s.MaxAccel, err = CalcAccel(y, distf)
	if err != nil {
		return err
	}

	fmt.Println(r, StatsHeader)
	fmt.Println(r, s.Pretty())
	return nil
}
