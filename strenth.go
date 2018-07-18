package main

import (
	"log"
	"math"
)

func strenth(f *force, sec *section) (isok bool) {

	epstion := math.Sqrt(235 / sec.fy())
	val := func(sec *section) int {
		is := sec.d / sec.thick
		switch {
		case is < 50*math.Pow(epstion, 2):
			return 0
		case is < 70*math.Pow(epstion, 2):
			return 1
		case is < 90*math.Pow(epstion, 2):
			return 2
		case is < 100*math.Pow(epstion, 2):
			return 3
		default:
			return -1
		}
	}(sec)

	gramma := func(v int) float64 {
		switch {
		case v > 2:
			return 1.0
		default:
			return 1.15
		}
	}(val)

	maxval := f.n/sec.area() + f.m/sec.wn()/gramma
	minval := f.n/sec.area() - f.m/sec.wn()/gramma
	log.Println("Max strenth is ", maxval)
	log.Println("Min strenth is ", minval)
	if maxval > sec.fy() || minval > sec.fy() {
		return false
	}
	return true
}
