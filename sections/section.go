package sections

import (
	"fmt"
	"math"
	"strings"
)

//Po stands for point data
type Po struct {
	ID, X, Y, Z int
}

func (s Po) String() string {
	return fmt.Sprintf("Point%v is ( %v,%v,%v)", s.ID, s.X, s.Y, s.Z)
}

//Body stand for the Tower frame
type Body struct {
	ID     int
	P1, P2 Po
	Section
}

//Section as it's name
type Section struct {
	D     float64
	Thick float64
	Grade string
}

//Wn return wn
func (s *Section) Wn() float64 {
	return math.Pi * math.Pow(s.D, 3) / 32.00
}

//Area return the area
func (s *Section) Area() float64 {
	return 0.25 * math.Pi * (math.Pow(s.D, 2) + 2*s.D*s.Thick - math.Pow(s.Thick, 2))
}

//Ix return ix
func (s *Section) Ix() float64 {
	return math.Pi * (math.Pow(s.D, 4) - math.Pow(s.D-s.Thick, 4)) / 64.0
}

//R return r
func (s *Section) R() float64 {
	return 0.354 * s.D
}

//Fy return fy value
func (s *Section) Fy() float64 {
	switch strings.ToUpper(s.Grade) {
	case "Q235":
		return 215
	case "Q345":
		return 310
	default:
		return 0
	}
}

//C return c
func (s *Section) C() float64 {
	return 0.5 * (s.D + s.Thick)
}
