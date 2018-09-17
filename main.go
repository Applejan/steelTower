//基于xls计算避雷针的强度和稳定
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	var height int
	var grade string
	bodys := make([]body, 50)
	points := make([]po, 50)

	//Init points and bodys
	xls, err := excelize.OpenFile("Model.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	rows := xls.GetRows("Input")
	for k, v := range rows {
		switch v[0] {
		case "Total_High":
			height, _ = strconv.Atoi(v[1])
			fallthrough
		case "Steel_Grade":
			grade = v[1]
			fallthrough
		case "Wind_w0":
			windInfo.w0, _ = strconv.ParseFloat(v[1], 64)
			fallthrough
		case "Ground_rou":
			windInfo.rou = rough(v[1])
			fallthrough
		case "Id":
			k++
			for j := 0; j < height; j++ {
				bodys[j].section.d, _ = strconv.ParseFloat(rows[k][1], 64)
				bodys[j].section.thick, _ = strconv.ParseFloat(rows[k][2], 64)
				k++
			}
		}
	}
	for i := 0; i < height+1; i++ {
		points[i].id = i + 1
		points[i].z = i
	}
	for i := 0; i < height; i++ {
		bodys[i].id = i + 1
		bodys[i].p1 = points[i]
		bodys[i].p2 = points[i+1]
		bodys[i].grade = grade

	}
	points = points[:height+1]
	bodys = bodys[:height]

	//Reset the xls file info
	var strs = []string{
		"Connectivity - Frame",
		"Frame Loads - Distributed",
		"Frame Props 01 - General",
		"Frame Section Assignments",
		"Joint Coordinates",
	}
	for _, v := range strs {
		rows := xls.GetRows(v)
		for i := 2; i < len(rows); i++ {
			xls.RemoveRow(v, 3)
		}
	}

	//Write Joint
	// xls.SetActiveSheet(xls.GetSheetIndex("Joint Coordinates"))
	log.Println("Now write Joint infomation")
	for i, v := range points {
		index := i + 4
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("A", index), v.id)
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("H", index), v.x)
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("E", index), v.y)
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("I", index), v.y)
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("D", index), v.x)
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("F", index), v.z)
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("J", index), v.z)
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("B", index), "GLOBAL")
		xls.SetCellValue("Joint Coordinates", fmt.Sprint("C", index), "Cartesian")
	}

	log.Println("Now write Frame infomation")
	for i, v := range bodys {
		index := i + 4
		xls.SetCellValue("Connectivity - Frame", fmt.Sprint("A", index), v.id)
		xls.SetCellValue("Connectivity - Frame", fmt.Sprint("B", index), v.p1.id)
		xls.SetCellValue("Connectivity - Frame", fmt.Sprint("C", index), v.p2.id)

		secName := fmt.Sprintf("Pipe%.0f*%.0f", v.section.d, v.section.thick)
		xls.SetCellValue("Frame Props 01 - General", fmt.Sprint("A", index), secName)
		xls.SetCellValue("Frame Props 01 - General", fmt.Sprint("B", index), v.grade)
		xls.SetCellValue("Frame Props 01 - General", fmt.Sprint("C", index), "Pipe")
		xls.SetCellValue("Frame Props 01 - General", fmt.Sprint("D", index), floatTofloat(v.d/1000.0, 3))
		xls.SetCellValue("Frame Props 01 - General", fmt.Sprint("E", index), floatTofloat(v.thick/1000.0, 3))

		xls.SetCellValue("Frame Section Assignments", fmt.Sprint("A", index), v.id)
		xls.SetCellValue("Frame Section Assignments", fmt.Sprint("B", index), "Pipe")
		xls.SetCellValue("Frame Section Assignments", fmt.Sprint("D", index), secName)

		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("A", index), v.id)
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("B", index), "WIND")
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("C", index), "GLOBAL")
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("D", index), "FORCE")
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("E", index), "X")
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("F", index), "RelDist")
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("G", index), 0)
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("H", index), 1)
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("K", index), floatTofloat(v.wind(), 3))
		xls.SetCellValue("Frame Loads - Distributed", fmt.Sprint("L", index), floatTofloat(v.wind(), 3))
	}

	xls.Save()
}

func floatTofloat(x float64, bit int) (val float64) {
	if bit == 2 {
		val, _ = strconv.ParseFloat(fmt.Sprintf("%0.2f", x), 64)
	} else if bit == 3 {
		val, _ = strconv.ParseFloat(fmt.Sprintf("%0.3f", x), 64)
	}
	return
}
