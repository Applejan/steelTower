package main

type po struct {
	id, x, y, z int
}

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

func (s *body) xi() (uz float64) {
	t := 0.07 * float64(s.height)
	tt := s.winddata.w0 * t * t
	num := [][]float64{
		{0.01, 1.47}, {0.02, 1.57}, {0.04, 1.69},
		{0.06, 1.77}, {0.08, 1.83}, {0.10, 1.88},
		{0.20, 2.04}, {0.40, 2.24}, {0.60, 2.36},
		{0.80, 2.46}, {1.00, 2.53}, {2.00, 2.80},
		{4.00, 3.09}, {6.00, 3.28}, {8.00, 3.42},
		{10.00, 3.54}, {20.00, 3.91}, {30.00, 4.14},
	}
	for i, v := range num {
		if v[0] > tt {
			return num[i+1][1]
		}
	}
	return
}

func (s *body) epsilon1() (eps float64) {
	val := [][]float64{
		{10, 0.57, 0.72, 1.03, 1.66},
		{20, 0.51, 0.63, 0.87, 1.35},
		{40, 0.45, 0.55, 0.73, 1.06},
		{80, 0.39, 0.46, 0.58, 0.80},
	}
	for _, v := range val {
		if int(v[0]) > s.height {
			return v[s.winddata.rou.id()+1]
		}
	}
	return
}

//epsilon2 is not work
func (s *body) epsilon2() (eps float64) {
	val := [][]float64{
		{10, 0.57, 0.72, 1.03, 1.66},
		{20, 0.51, 0.63, 0.87, 1.35},
		{40, 0.45, 0.55, 0.73, 1.06},
		{80, 0.39, 0.46, 0.58, 0.80},
	}
	for _, v := range val {
		if int(v[0]) > s.height {
			return v[s.winddata.rou.id()+1]
		}
	}
	return
}

func (s *body) wind() float64 {
	beta := 1 + s.xi()*s.epsilon1()*s.epsilon2()
	return 0.9 * s.uz() * beta * s.winddata.w0
}
