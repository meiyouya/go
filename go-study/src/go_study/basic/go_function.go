package main

import "math"

func sum(num1, num2 int) int {
	return num1 + num2
}

func reverse(x, y string) (string, string) {
	return y, x
}

//func swap(x, y int)  {
//	temp := x
//	x = y
//	y = temp
//}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**
定义一个结构体
*/
type Circle struct {
	radius float64
}

/**
定义属于Circle的方法
*/
func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var c Circle
	c.radius = 10.00
	println("圆的面积为：", c.getArea())

	//nextNumber := getSequence()		// 每次调用getSequence()时，i被重置为0
	//
	//println(nextNumber())	// 每次调用nextNumber()，i会在上一次调用nexNumber()的i值上加1
	//println(nextNumber())
	//println(nextNumber())
	//
	//nextNumber1 := getSequence()
	//println(nextNumber1())
	//println(nextNumber1())
	//println(nextNumber1())

	//getSquare := func(x float64) float64 {
	//	return math.Sqrt(x)
	//}
	//fmt.Println(getSquare(9)) // 输出 3

	//x, y := 3, 4
	//fmt.Println("交换之前x、y的值：",x,y)	// 输出 交换之前x、y的值： 3 4
	//swap(&x,&y)
	//fmt.Println("交换之后x、y的值：",x,y)	// 输出 交换之前x、y的值： 4 3

	//x, y := 3, 4
	//fmt.Println("交换之前x、y的值：",x,y)	// 输出 交换之前x、y的值： 3 4
	//swap(x,y)
	//fmt.Println("交换之后x、y的值：",x,y)	// 输出 交换之前x、y的值： 3 4

	//i := sum(3, 5) // 调用函数并返回值
	//fmt.Print(i)   // 输出 8
	//a, b := reverse("haha", "xixi")
	//fmt.Println(a, b) // 输出  xixi    haha
}
