package main

import "fmt"

func main() {
	var arr1 [3]int // 声明数组

	arr1[0] = 1
	arr1[1] = 3
	arr1[2] = 5
	fmt.Println(arr1) // 按照索引初始化

	var arr2 = [5]int{1, 3, 5, 7}
	fmt.Println(arr2) // 在声明的时候就初始化，[]指定长度，{}中是初始化的值，个数不能大于长度，不足的补该类型默认值，如int默认0

	var arr3 = []int{1, 3, 4, 5, 6, 7, 8, 89, 2}
	fmt.Println(arr3) // 在声明的时候就初始化，[]为空，不指定长度，{}中是初始化的值，长度根据实际的个数自动设置

	fmt.Println(arr1[2]) // 访问数组arr1中的索引为2的元素

	var arr4 [3][4]int // 声明一个二维数组

	arr4[0][0] = 1
	arr4[0][1] = 2
	arr4[0][2] = 3
	arr4[0][3] = 4
	arr4[1][0] = 5
	arr4[1][1] = 6
	arr4[1][2] = 7
	arr4[1][3] = 8
	arr4[2][0] = 9
	arr4[2][1] = 10
	arr4[2][2] = 11
	arr4[2][3] = 12
	fmt.Println(arr4) // 声明后通过索引初始化

	var arr5 = [3][4]int{
		{1, 3, 5, 7},
		{2, 4, 6, 8},
		{3, 6, 9, 0},
	}
	fmt.Println(arr5) // 声明的时候就初始化，指定各维长度，不指定的维度根据实际长度自动设置

	fmt.Println(arr5[1][2]) // 通过索引访问值
}
