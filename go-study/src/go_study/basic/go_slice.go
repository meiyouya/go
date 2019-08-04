package main

import "fmt"

func main() {
	var slice1 []int               // 不声明长度的切片，这总定义方式的切片，在初始化之前为nil
	var slice2 = make([]int, 3, 5) // 通过make函数创建切片，第三个参数capacity，用于设置容量，非必填
	slice3 := make([]int, 3)       // 简化写法

	fmt.Println(slice1 == nil) // 输出 true

	// 初始化切片
	slice1 = []int{1, 2, 3}
	slice2[0] = 4
	slice2[1] = 5
	slice2[2] = 6
	slice3[0] = 7
	slice3[1] = 8
	slice3[2] = 9

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	fmt.Println(len(slice1)) // 输出3
	fmt.Println(cap(slice2)) // 输出5

	numbers := slice3[:1] // Go的切片可以被截取，类似于Python的列表切片
	fmt.Println(numbers)  // 输出  [7]

	// append函数用于向切片增加元素
	slice3 = append(slice3, 8)
	fmt.Println(slice3)

	// 如果需要一个更大容量的切片，又不想损失原来的内容，就可以通过copy函数复制
	slice4 := make([]int, 3, 10)
	copy(slice4, slice2)
	fmt.Println(slice4)
}
