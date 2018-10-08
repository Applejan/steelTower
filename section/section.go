package section

import (
	"fmt"
	"math"
	"strings"
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
	force float64
}


type section struct {
	d     float64
	thick float64
	grade string
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

func (s *section) fy() float64 {
	switch strings.ToUpper(s.grade) {
	case "Q235":
		return 215
	case "Q345":
		return 310
	default:
		return 0
	}
}
func (s *section) c() float64 {
	return 0.5 * (s.d + s.thick)
}
