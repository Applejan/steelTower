package main

import (
	"fmt"
	"math"
)

//GB50017-2017强度验算
func strenth1(f *force, sec *section) (isok bool) {
	epstion := math.Sqrt(235 / sec.Fy())
	val := func(sec *section) int {
		is := sec.D / sec.Thick
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

	maxval := f.n/sec.Area() + f.m/sec.Wn()/gramma
	minval := f.n/sec.Area() - f.m/sec.Wn()/gramma
	if maxval > sec.Fy() || minval > sec.Fy() {
		fmt.Printf("Body%v in %v force is M=%.2f,V=%.2f,N=%.2f(GB50017)\n", f.frameID, f.forceID, f.m, f.v, f.n)
		return false
	}
	return true
}

//DLT5130-2001弯曲强度计算
func strenth2(s *section, f *force) (isok bool) {
	if f.m*s.C()/s.Ix() <= fb(s) {
		return true
	}
	fmt.Printf("Body%v in %v force is M=%.2f,V=%.2f,N=%.2f(DLT5130),Want %v \n", f.frameID, f.forceID, f.m, f.v, f.n, fb(s))
	return
}

//DLT5130-2001剪切强度计算
func strenth3(s *section, f *force) (isok bool) {
	if f.v*qit(s) <= 0.58*s.Fy() {
		return true
	}
	fmt.Printf("Body%v in %v force is M=%.2f,V=%.2f,N=%.2f(DLT5130),Want %v \n", f.frameID, f.forceID, f.m, f.v, f.n, 0.58*s.Fy())
	return

}

//DLT5130-2001复合受力强度计算
func strenth4(s *section, f *force) (isok bool) {
	tmp1 := math.Pow(f.n/s.Area()+f.m*s.C()/s.Ix(), 2)
	tmp2 := 3 * math.Pow(f.v*qit(s), 2)
	if tmp1*tmp1+tmp2*tmp2 <= math.Pow(fb(s), 2) {
		return true
	}
	fmt.Printf("Body%v in %v force is M=%.2f,V=%.2f,N=%.2f(DLT5130)\n", f.frameID, f.forceID, f.m, f.v, f.n)
	return
}

// Strength implments the check of strength
func Strength(sections map[string]section, f *force) {
	s := sections[f.frameID]
	_ = strenth4(&s, f)
	_ = strenth2(&s, f)
	_ = strenth3(&s, f)
	_ = strenth4(&s, f)
}
