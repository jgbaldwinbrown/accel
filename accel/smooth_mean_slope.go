package accel

import (
	"sort"
	"github.com/chewxy/stl/loess"
	"github.com/montanaflynn/stats"
)

type Af struct {
	Pop string
	Gen int
	Af float64
}

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

func MeanAfs(afs []Af) (means []float64) {
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

func SortedGens(afs []Af) (gens []float64) {
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

func SmoothedMeanSlope(afs []Af) ([]float64, []float64) {
	means := MeanAfs(afs)
	smeans := Smooth(means)
	gens := SortedGens(afs)
	splits, slopes, err := Deriv(gens, smeans)
	if err != nil {
		panic(err)
	}
	return splits, slopes
}
