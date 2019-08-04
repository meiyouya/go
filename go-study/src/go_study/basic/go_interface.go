package main

import "fmt"

// 定义一个动物接口
type Animal interface {
	eat()
	run() bool
}

type Cat struct {
}

func (cat Cat) eat() {
	fmt.Println("我是一只猫")
}
func (cat Cat) run() bool {
	return false
}

type Dog struct {
}

func (dog Dog) eat() {
	fmt.Println("我是一只狗")
}

func (dog Dog) run() bool {
	return true
}

func main() {
	var animal Animal

	animal = new(Cat)
	animal.eat()              // 输出    我是一只猫
	fmt.Println(animal.run()) // 输出    false

	animal = new(Dog)
	animal.eat()              // 输出    我是一只狗
	fmt.Println(animal.run()) // 输出    true
}
