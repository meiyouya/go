package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s { // 第一个_实际上是索引，但是索引用不到，直接用变量声明就必须要使用，用 _  代表占个位
		sum += v
	}
	ch <- sum // 把 sum  发给通道ch
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 如果这里不关闭，那么从 c 这个通道中接收数据接收到第11个的时候就会被阻塞
}

func main() {
	//go say("world")
	//say("hello")

	/*s := []int{2,4,6,8,9,2}
	c := make(chan int)		// 通道在使用之前必须先创建
	go sum(s[:len(s)/2],c)		// 计算后一半数据的和
	go sum(s[len(s)/2:],c)		// 计算前一半数据的和

	x,y := <- c, <- c		// 从通道中接收数据

	fmt.Println(x,y,x+y)*/

	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
