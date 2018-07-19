package main

type po struct {
	id, x, y, z int
}

type body struct {
	id     int
	p1, p2 int
	section
	wind
	force float64
}
type wind struct{}

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

// func wind(s *section, ro rough) (w float64) {
// id:=ro.id()
// f:=1/0.07/height

//风荷载计算相关表格
// 	phi := [][]float64{
// 		{0.1, 0.02},
// 		{0.2, 0.06},
// 		{0.3, 0.14},
// 		{0.4, 0.23},
// 		{0.5, 0.34},
// 		{0.6, 0.46},
// 		{0.7, 0.59},
// 		{0.8, 0.79},
// 		{0.9, 0.86},
// 		{1.0, 1.00},
// 	}
// 	theta:=[]
// }
