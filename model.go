package main

type po struct {
	id, x, y, z int
}

type body struct {
	id     int
	p1, p2 int
	section
	windload func() int
	force    float64
}

type force struct {
	id string
	m  float64
	v  float64
	n  float64
}
