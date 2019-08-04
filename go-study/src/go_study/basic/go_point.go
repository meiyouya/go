package main

import "fmt"

func main() {
	/*var a int = 123		// 定义一个变量
	var ip *int = &a		// 声明一个指针
	fmt.Println(a)		// 输出 123
	fmt.Println(&a)		// 输出 0xc000054080
	fmt.Println(ip)		// 输出 0xc000054080
	fmt.Println(*ip)	// 输出 123*/

	/*arr := []int{10,100,200}	// 定义一个数组
	var i int
	var ptr [3]*int		// 定义一个指针数组

	for i = 0; i < 3; i++ {
		ptr[i] = &arr[i]	// 取arr每个元素的地址值
	}

	for i = 0; i < 3; i++ {
		fmt.Printf("arr[%d]的指针：%d\n",i,ptr[i])		// 输出每个指针的值
		fmt.Printf("arr[%d] = %d\n",i,*ptr[i])		// 输出每个指针指向的值
	}*/

	/*var a int	// 声明一个整数
	var ptr *int	// 声明一个指针
	var pptr ** int		// 声明一个指向指针的指针

	a = 3000
	ptr = &a
	pptr = &ptr

	fmt.Println("a的值：",a)
	fmt.Println("ptr的值：",ptr)
	fmt.Println("ptr指向的值：",*ptr)
	fmt.Println("pptr的值：",pptr)
	fmt.Println("pptr指向的值：",*pptr)*/

	var a = 3
	var b = 4
	var c = 5
	add(&a, &b, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// 声明一个接收指针类型参数的函数
func add(x *int, y *int, z int) {
	fmt.Println("x的值：", x)
	fmt.Println("x指向的值：", *x)
	fmt.Println("y的值：", y)
	fmt.Println("y指向的值：", *y)

	*x = 7 // 直接修改值，由于传递进来的指针，所以a的实际值会跟着变
	z = 9  // 非指针类型是通过值传递，此处修改不会影响到c的值
}
