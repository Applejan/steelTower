package main

import (
	"log"
)

func strenth(f *force, sec *section) (isok bool) {
	maxval := f.n/sec.area() + f.m/sec.wn()
	minval := f.n/sec.area() - f.m/sec.wn()
	log.Println("Max strenth is ", maxval)
	log.Println("Min strenth is ", minval)
	if maxval > sec.f || minval > sec.f {
		return false
	}
	return true
}
