//基于xls计算避雷针的强度和稳定
package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strconv"
)

func main() {
	var height int
	var grade string
	bodys := make([]body, 50)
	points := make([]po, 50)
	//Init points and bodys
	xls, err := excelize.OpenFile("Init_file.xlsx")
	if err != nil {
		log.Println(err)
	}
	rows := xls.GetRows("Sheet1")
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
			for j := k; j < height+1; j++ {
				bodys[j].section.d, _ = strconv.ParseFloat(v[1], 64)
				bodys[j].section.thick, _ = strconv.ParseFloat(v[2], 64)
			}
		}
	}
	for i := 0; i < height+1; i++ {
		points[i].id = i
		points[i].x = 0
		points[i].y = 0
		points[i].z = i
	}
	for i := 0; i < height; i++ {
		bodys[i].id = i
		bodys[i].p1 = points[i]
		bodys[i].p2 = points[i+1]
		bodys[i].height = height
		bodys[i].section.degree = grade
	}
	points = points[:height+1]
	bodys = bodys[:height]
	fmt.Println(points)

}
