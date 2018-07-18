package main

import (
	"math"
)

type section struct {
	d     float64
	thick float64
	f     float64
}

func (s *section) wn() float64 {
	return math.Pi * math.Pow(s.d, 3) / 32.00
}

func (s *section) area() float64 {
	return 0.25 * math.Pi * (math.Pow(s.d, 2) + 2*s.d*s.thick - math.Pow(s.thick, 2))
}

func (s *section) ix() float64 {
	return math.Pi * (math.Pow(s.d, 4) - math.Pow(s.d-s.thick, 4)) / 64.0
}

func (s *section) r() float64 {
	return 0.354 * s.d
}
