package main

import (
	"fmt"
	"math"
)

//Point data
type po struct {
	id, x, y, z int
}

func (s *po) String() string {
	return fmt.Sprintf("Point%v is ( %v,%v,%v)", s.id, s.x, s.y, s.z)
}

//body stand for the Tower frame
type body struct {
	id     int
	p1, p2 po
	section
	winddata
	force  float64
	height int
}

type winddata struct {
	w0  float64
	rou rough
}

type force struct {
	id string
	m  float64
	v  float64
	n  float64
}

type rough string

var rou rough

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
	uzs := [][]float64{
		{5, 1.17, 1.00, 0.74, 0.62},
		{10, 1.38, 1.00, 0.74, 0.62},
		{15, 1.52, 1.14, 0.74, 0.62},
		{20, 1.63, 1.25, 0.84, 0.62},
		{30, 1.80, 1.42, 1.00, 0.62},
		{40, 1.92, 1.56, 1.13, 0.73},
		{50, 2.03, 1.67, 1.25, 0.84},
	}
	i := s.winddata.rou.id()
	for k, v := range uzs {
		if int(v[k]) < s.p1.z {
			tmp1 := uzs[k][i+1]
			tmp2 := uzs[k+1][i+1]
			return tmp1 + (tmp2-tmp1)/(uzs[k+1][0]-uzs[k][i+1])*(float64(s.p1.z)-uzs[k][i+1])
		}
	}
	return
}

func (s *body) wind() float64 {
	us := func() (us float64) {
		tmp := s.uz() * s.winddata.w0 * math.Pow(s.section.d, 2)
		if tmp <= 0.002 {
			return 1.2
		} else if tmp >= 0.15 {
			return 0.6
		}
		return 1.2 + (tmp-0.002)*0.6/0.013
	}()
	beta := 2.0 //风振系数
	return 0.9 * s.uz() * beta * us * s.winddata.w0
}
