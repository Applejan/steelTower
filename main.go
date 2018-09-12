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
			for j, _ := strconv.Atoi(rows[k][0]); j < height+1; j++ {
				bodys[j].section.d, _ = strconv.ParseFloat(rows[k][1], 64)
				bodys[j].section.thick, _ = strconv.ParseFloat(rows[k][2], 64)
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
		bodys[i].windForce = bodys[i].wind()
	}
	points = points[:height+1]
	bodys = bodys[:height]

	for _, v := range bodys {

		fmt.Println(v)
	}

}
