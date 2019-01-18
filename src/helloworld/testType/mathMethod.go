package testType

import "fmt"

//求圆的面积
const PI float32 = 3.1415926

func Circle(r float32) {
	var area = PI * r * r
	fmt.Println("换算的面积是：", area)
}
