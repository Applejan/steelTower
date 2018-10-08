package main

func main() {
	frame := new(body)
	//init the frame
	//use goroute to check stength and stablity

	Strength(&frame.section, &frame.force)
	Stablity(&frame.section, &frame.force)
}
