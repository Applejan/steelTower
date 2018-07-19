package main

import (
	"fmt"
	"log"
	"math"
)

type sections interface {
	wn() float64
	area() float64
	ix() float64
	r() float64
	fy() float64
}

//GB50017-2017强度验算
func strenth1(f *force, sec *section) (isok bool) {

	epstion := math.Sqrt(235 / sec.fy())
	val := func(sec *section) int {
		is := sec.d / sec.thick
		switch {
		case is <= 50*math.Pow(epstion, 2):
			return 0
		case is <= 70*math.Pow(epstion, 2):
			return 1
		case is <= 90*math.Pow(epstion, 2):
			return 2
		case is <= 100*math.Pow(epstion, 2):
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

//DLT5130-2001弯曲强度计算
func strenth2(s *section, f *force) (isok bool) {
	if f.m*s.c()/s.ix() <= fb(s) {
		return true
	}
	return
}

//DLT5130-2001剪切强度计算
func strenth3(s *section, f *force) (isok bool) {
	if f.v*qit(s) <= 0.58*s.fy() {
		return true
	}
	return
}

//DLT5130-2001复合受力强度计算
func strenth4(s *section, f *force) (isok bool) {
	tmp1 := math.Pow(f.n/s.area()+f.m*s.c()/s.ix(), 2)
	tmp2 := 3 * math.Pow(f.v*qit(s), 2)
	if tmp1*tmp1+tmp2*tmp2 <= math.Pow(fb(s), 2) {
		return true
	}
	return
}

// Strength implments the check of strength
func Strength(s *section, f *force) {
	if strenth4(s, f) {
		fmt.Println("GB50017-2017强度验算,OK!")
	} else {
		fmt.Println("GB50017-2017强度验算,False!")
	}
	if strenth2(s, f) {
		fmt.Println("DLT5130-2001弯曲强度计算,OK!")
	} else {
		fmt.Println("DLT5130-2001弯曲强度计算,False!")
	}
	if strenth3(s, f) {
		fmt.Println("DLT5130-2001剪切强度计算,OK!")
	} else {
		fmt.Println("DLT5130-2001剪切强度计算,False!")
	}
	if strenth4(s, f) {
		fmt.Println("DLT5130-2001剪切强度计算,OK!")
	} else {
		fmt.Println("DLT5130-2001剪切强度计算,False!")
	}
}
