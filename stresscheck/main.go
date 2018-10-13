package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/applejan/steelTower/sections"
	"log"
	"os"
	"strconv"
	"strings"
)

//force like it called
type force struct {
	frameID, forceID string
	m                float64
	v                float64
	n                float64
}

type section struct {
	sections.Section
}

func main() {
	grade := "Q235"
	fileName := os.Args[1]
	xls, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Println(err)
	}
	rows := xls.GetRows("Frame Section Assignments")
	//init the frame
	frames := make(map[string]section)
	forces := make([]force, len(rows)-3)
	for i, v := range rows {
		if i < 3 {
			continue
		}
		id := v[0]
		d, thick := func(s string) (d float64, thick float64) {
			s = strings.TrimPrefix(s, "Pipe")
			tmp := strings.Split(s, "*")
			d, _ = strconv.ParseFloat(tmp[0], 64)
			thick, _ = strconv.ParseFloat(tmp[1], 64)
			return
		}(v[3])

		frames[id] = section{sections.Section{D: d, Thick: thick, Grade: grade}}
	}

	rows = xls.GetRows("Element Forces - Frames")
	for index, val := range rows {
		if index < 3 {
			continue
		}
		if val[4] == "Mode" {
			continue
		}
		frameID := val[0]
		forceID := val[2]
		m, _ := strconv.ParseFloat(val[11], 64)
		v, _ := strconv.ParseFloat(val[7], 64)
		p, _ := strconv.ParseFloat(val[6], 64)
		ff := force{frameID, forceID, m, v, p}
		forces = append(forces, ff)
		checkIt(frames, forces)
	}
}

func checkIt(frame map[string]section, f []force) {
	for _, ff := range f {
		Strength(frame, &ff)
	}
}
