package accel

import (
	"sort"
	"github.com/chewxy/stl/loess"
	"github.com/montanaflynn/stats"
	"math"
)

type Af struct {
	Pop string
	Gen int
	Af float64
}

type Afs []Af

type IntSet struct {
	List []int
	Map map[int]struct{}
}

func NewIntSet() *IntSet {
	i := new(IntSet)
	i.Map = make(map[int]struct{})
	return i
}

func (s *IntSet) Add(i int) {
	_, ok := s.Map[i]
	if !ok {
		s.Map[i] = struct{}{}
		s.List = append(s.List, i)
	}
}

func MeanAfs(afs Afs) (means []float64) {
	gens := NewIntSet()
	for _, af := range afs {
		gens.Add(af.Gen)
	}
	for _, gen := range gens.List {
		var tomean []float64
		for _, af := range afs {
			if af.Gen == gen {
				tomean = append(tomean, af.Af)
			}
		}
		mean, err := stats.Mean(tomean)
		if err != nil {
			panic(err)
		}
		means = append(means, mean)
	}
	return means
}

func Smooth(fs []float64) []float64 {
	y, err := loess.Smooth(fs, 3000, 1, loess.Linear)
	if err != nil {
		panic(err)
	}
	return y
}

func SortedGens(afs Afs) (gens []float64) {
	genset := NewIntSet()
	for _, af := range afs {
		genset.Add(af.Gen)
	}

	for _, gen := range genset.List {
		gens = append(gens, float64(gen))
	}
	sort.Float64s(gens)
	return gens
}

func SmoothedMeanSlope(afs Afs) ([]float64, []float64, error) {
	means := MeanAfs(afs)
	smeans := Smooth(means)
	gens := SortedGens(afs)
	splits, slopes, err := Deriv(gens, smeans)
	if err != nil {
		return nil, nil, err
	}
	return splits, slopes, nil
}

func MaxIndex(fs []float64) int {
	max := math.Inf(-1)
	maxi := -1
	for i, f := range fs {
		if f > max {
			max = f
			maxi = i
		}
	}
	return maxi
}

func MaxSlopeTimes(afsets []Afs) ([]float64, error) {
	out := make([]float64, len(afsets))
	for i, afs := range afsets {
		xs, ys, err := SmoothedMeanSlope(afs)
		if err != nil {
			return nil, err
		}
		out[i] = xs[MaxIndex(ys)]
	}
	return out, nil
}
