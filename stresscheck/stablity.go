package main

import (
	"errors"
	"fmt"
)

//DLT5130中受压稳定强度设计值
func fc(s *section) (fc float64) {
	k := s.D / s.Thick
	if k <= 24100/s.Fy() {
		return s.Fy()
	}
	if k > 24100/s.Fy() && k <= 76130/s.Fy() {
		return 0.75*s.Fy() + 6025/k
	}
	return
}

//DLT5130  Q/It  最大弯曲剪应力系数
func qit(s *section) (fc float64) {
	return 0.637 / s.D / s.Thick
}

//DLT5130中受弯稳定强度设计值
func fb(s *section) (fc float64) {
	k := s.D / s.Thick
	if k <= 38060/s.Fy() {
		return s.Fy()
	}
	if k > 38060/s.Fy() && k <= 76130/s.Fy() {
		return 0.7*s.Fy() + 11410/k
	}
	return
}

//DLT5130-2001局部稳定验算
func stablity(s *section, f *force) (percent float64, err error) {
	v := -f.n/s.Area()/fc(s) + f.m*s.C()/s.Ix()/fb(s)
	if v <= 1 {
		return v, err
	}
	str := fmt.Sprintf("Body%v in %v force is M=%.2f,V=%.2f,N=%.2f(DLT5130)\n", f.frameID, f.forceID, f.m, f.v, f.n)
	return v, errors.New(str)
}

//Stablity implents the check of stablity
func Stablity(section map[string]section, f *force) (float64, error) {
	s := section[f.frameID]
	str, err := stablity(&s, f)
	return str, err
}
