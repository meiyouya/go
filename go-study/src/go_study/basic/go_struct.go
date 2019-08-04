package main

import "fmt"

// 定义一个名称为Books的结构体
type Books struct {
	id     int    // 书的编号，类型为int
	title  string // 书的标题，类型为string
	author string // 书的作者，类型为string
}

func main() {
	// 结构体有点类似Java中的对象，创建新的结构体类似于对象的实例化
	// 按照结构体中的字段顺序赋值
	goBook := Books{1, "Go语言从入门到精通", "Go语言大佬"}
	// 通过字段的名称赋值，这时候顺序就不重要了
	cBook := Books{id: 2, author: "C语言大佬", title: "C语言从入门到精通"}
	// 没有赋值的字段，默认为该类型的空值
	pythonBook := Books{author: "Python语言大佬"}
	fmt.Println(goBook)     // 输出   {1 Go语言从入门到精通 Go语言大佬}
	fmt.Println(cBook)      // 输出   {2 C语言从入门到精通 C语言大佬}
	fmt.Println(pythonBook) // 输出   {0  Python语言大佬}

	fmt.Printf(goBook.title) // 访问结构体成员，通过  点  语法访问

	printBook(goBook)
	fmt.Println(goBook)
	printBookByPoint(&cBook)
	fmt.Println(cBook)
}

// 结构体也可以作为函数的参数进行传递
func printBook(book Books) {
	fmt.Println(book.id)
	fmt.Println(book.title)
	fmt.Println(book.author)

	book.title = "Go语言改一改" // 这里修改值不会影响到goBook
}

// 结构体也可以使用指针传递
func printBookByPoint(book *Books) {
	fmt.Println(book.id)
	fmt.Println(book.title)
	fmt.Println(book.author)

	book.title = "C语言改一改" // 这里修改值会影响到cBook
}
