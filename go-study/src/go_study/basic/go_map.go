package main

import "fmt"

func main() {
	// 定义map有2种方式
	var map1 map[string]string     // 这种方式定义，是不能被赋值的，因为其没有被初始化，默认为nil map
	map1 = make(map[string]string) // 定义之后必须初始化才可以使用
	map2 := make(map[string]string)

	map1["apple"] = "苹果"
	map1["banana"] = "香蕉"
	map2["orange"] = "橘子"
	map2["pairs"] = "梨子"

	for fruit := range map1 {
		fmt.Println(fruit)
	}
	for fruit := range map2 {
		fmt.Println(fruit)
	}

	capital, ok := map1["apple"] // 判断该元素是否存在，返回的第一个参数是该元素的值，第二个参数是是否存在的bool类型的值
	fmt.Println(capital)         // 输出  苹果
	fmt.Println(ok)              // 输出true

	delete(map2, "pairs") // 删除指定集合中的指定元素
	for fruit := range map2 {
		fmt.Println(fruit)
	}
}
