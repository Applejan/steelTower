package main

import (
	"fmt"
)

//Point data
type po struct {
	id, x, y, z int
}

func (s po) String() string {
	return fmt.Sprintf("Point%v is ( %v,%v,%v)", s.id, s.x, s.y, s.z)
}

//body stand for the Tower frame
type body struct {
	id     int
	p1, p2 po
	section
	windForce float64
	force     float64
	height    int
}

type force struct {
	id string
	m  float64
	v  float64
	n  float64
}
