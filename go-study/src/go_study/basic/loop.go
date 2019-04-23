package main

import "fmt"

func main() {

	for a := 3; a < 10; a++ {
		fmt.Print(a, ",")
	}

	fmt.Println()

	var a = 5
	for a < 10 {
		a++
		fmt.Print(a, ",")
	}

	for {
		// 这里的语句会进入无限循环
		fmt.Print(1)
	}
}
