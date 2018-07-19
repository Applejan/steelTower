//基于xls计算避雷针的强度和稳定
package main

func main() {
	//Init points and bodys
	var height int
	var bodys []body
	var points []po
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
	}

	//TOdo

	return
}
