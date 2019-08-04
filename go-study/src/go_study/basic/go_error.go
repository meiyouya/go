package main

import (
	"errors"
	"fmt"
	"math"
)

// 函数返回一个错误，类似于Java中的抛出异常
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("参数不能小于0")
	}
	return math.Sqrt(f), nil
}

// 自定义error的结构体，用于存储错误信息
type DivideError struct {
	dividee int
	divider int
}

// 自定义error，实现error接口
func (de *DivideError) Error() string {
	strFormat := `
		0不能做除数.
		dividee: %d
		divide: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(dividee int, divider int) (result int, errorMsg string) {
	if divider == 0 {
		data := DivideError{dividee: dividee, divider: divider}
		errorMsg = data.Error()
		return
	} else {
		return dividee / divider, ""
	}
}

func main() {
	result, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	r, errorMsg := Divide(100, 0)
	if errorMsg == "" {
		fmt.Println(r)
	} else {
		fmt.Println(errorMsg)
	}
}
