package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	sum := 0

	for index, num := range numbers {
		fmt.Println(index)
		sum += num
	}
	fmt.Println(sum)

	fruits := map[string]string{"apple": "苹果", "banana": "香蕉"}
	for key, value := range fruits {
		fmt.Printf("%s--%s\n", key, value)
	}

	for index, unicodeValue := range "go" {
		fmt.Println(index, unicodeValue)
	}
}
