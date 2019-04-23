package main

import (
	"fmt"
	"math"
)

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

func main() {
	getSquare := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquare(9)) // 输出 3

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
