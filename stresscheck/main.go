package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/applejan/steelTower/sections"
	"log"
	"os"
	"strconv"
	"strings"
)

//force like it called
type force struct {
	id string
	m  float64
	v  float64
	n  float64
}

type section struct {
	sections.Section
	f []force
}

func main() {
	fileName := os.Args[1]
	xls, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Println(err)
	}
	rows := xls.GetRows("Frame Section Assignments")
	//init the frame
	frames := make(map[int]section)
	for i, v := range rows {
		if i < 3 {
			continue
		}
		id, err := strconv.Atoi(v[0])
		if err != nil {
			fmt.Println(err)
		}
		d, thick := func(s string) (d float64, thick float64) {
			s = strings.TrimPrefix(s, "Pipe")
			tmp := strings.Split(s, "*")
			d, _ = strconv.ParseFloat(tmp[0], 64)
			thick, _ = strconv.ParseFloat(tmp[1], 64)
			return
		}(v[3])

		frames[id] = section{sections.Section{D: d, Thick: thick}, f: make([]force, 9)}
	}

	rows = xls.GetRows("Element Forces - Frames")
	for i, v := range rows {
		if i < 3 {
			continue
		}

	}
}
