package main

import (
	"math"
)

type winddata struct {
	w0  float64
	rou rough
}
type rough string

var windInfo winddata

func (r rough) id() (id int) {
	switch r {
	case "A", "a":
		return 0
	case "B", "b":
		return 1
	case "C", "c":
		return 2
	case "D", "d":
		return 3
	default:
		return -1
	}
}

func (s *body) uz() (uz float64) {
	z := s.p2.z
	uzs := [][]float64{
		{5, 1.17, 1.00, 0.74, 0.62},
		{10, 1.38, 1.00, 0.74, 0.62},
		{15, 1.52, 1.14, 0.74, 0.62},
		{20, 1.63, 1.25, 0.84, 0.62},
		{30, 1.80, 1.42, 1.00, 0.62},
		{40, 1.92, 1.56, 1.13, 0.73},
		{50, 2.03, 1.67, 1.25, 0.84},
	}
	i := windInfo.rou.id() + 1
	if z < int(uzs[0][0]) {
		return uzs[0][i]
	}
	for index, v := range uzs {
		if z < int(v[0]) {
			tmp1 := uzs[index-1][i]
			tmp2 := uzs[index][i]
			return tmp1 + (tmp2-tmp1)/(uzs[index][0]-uzs[index-1][0])*(float64(z)-uzs[index-1][0])

		}
	}
	return
}

func (s *body) wind() float64 {
	us := func() (us float64) {
		tmp := s.uz() * windInfo.w0 * math.Pow(s.section.d, 2)
		if tmp <= 0.002 {
			return 1.2
		} else if tmp >= 0.15 {
			return 0.6
		}
		return 1.2 + (tmp-0.002)*0.6/0.013
	}()
	beta := 2.0 //风振系数
	return 0.9 * s.uz() * beta * us * windInfo.w0 * s.d / 1000.0
}
