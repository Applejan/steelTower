package main

import (
	"fmt"
	"log"
)

//DLT5130中受压稳定强度设计值
func fc(s *section) (fc float64) {
	k := s.d / s.thick
	if k <= 24100/s.fy() {
		return s.fy()
	}
	if k > 24100/s.fy() && k <= 76130/s.fy() {
		return 0.75*s.fy() + 6025/k
	}
	return
}

//DLT5130  Q/It  最大弯曲剪应力系数
func qit(s *section) (fc float64) {
	return 0.637 / s.d / s.thick
}

//DLT5130中受弯稳定强度设计值
func fb(s *section) (fc float64) {
	k := s.d / s.thick
	if k <= 38060/s.fy() {
		return s.fy()
	}
	if k > 38060/s.fy() && k <= 76130/s.fy() {
		return 0.7*s.fy() + 11410/k
	}
	return
}

//DLT5130-2001局部稳定验算
func stablity(s *section, f *force) (isok bool) {
	v := f.n/s.area()/fc(s) + f.m*s.c()/s.ix()/fb(s)
	log.Println("Local stability val is ", v)
	if v <= 1 {
		return true
	}
	return
}

//Stablity implents the check of stablity
func Stablity(s *section, f *force) {
	if stablity(s, f) {
		fmt.Println("DLT5130-2001局部稳定验算,OK!")
	} else {
		fmt.Println("DLT5130-2001局部稳定验算,False!")
	}
}
