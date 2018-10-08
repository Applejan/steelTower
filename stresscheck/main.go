package main
//force like it called
type force struct {
	id string
	m  float64
	v  float64
	n  float64
}

func main() {
	frame := new(body)
	//init the frame
	//use goroute to check stength and stablity

	Strength(&frame.section, &frame.force)
	Stablity(&frame.section, &frame.force)
}
