package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	//testAdd()
	//indefiniteParameter(1, 2, 3)
	//indefiniteParameterAnyType(1, "hello world", float32(99))
	//anonymous()
	//closure()

	//err := errorProcess("https://www.baidu.com")
	//if err != nil {
	//	fmt.Println("path error !!!", err)
	//}

	deferdemo()
}

// defer
func deferdemo() {
	defer func() {
		fmt.Println("1111111")
	}()

	defer fmt.Println("2222222")
	defer fmt.Println("3333333")

	for i := 0; i < 100000; i++ {
		// do nothing
	}

	return
}

// 错误处理
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (pathError PathError) Error() string {
	return pathError.Op + " " + pathError.Path + " : " + pathError.Err.Error()
}

func errorProcess(path string) (err error) {
	if !strings.HasPrefix(path, "http://") {
		err = PathError{"GET", path, errors.New("path is not valid")}
		return
	}

	return nil
}

// 闭包
func closure() {
	var j int = 5
	a := func() func() {
		i := 1
		return func() {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}

	x := a()
	x()
	j *= 2
	x()
}

// 匿名函数
func anonymous() {
	f := func(input string) {
		fmt.Println(input)
	}

	f("hehehe")

	(func(input string) {
		fmt.Println(input)
	})("hahaha")
}

// 不定参数
func indefiniteParameter(args ...int) {
	indefiniteParameterCount(args...)
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func indefiniteParameterCount(args ...int) {
	fmt.Printf("共输入%d个参数\n", len(args))
}

// 任意类型的不定参数
func indefiniteParameterAnyType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is int")
		case int8:
			fmt.Println(arg, "is int8")
		case int16:
			fmt.Println(arg, "is int16")
		case int32:
			fmt.Println(arg, "is int32")
		case int64:
			fmt.Println(arg, "is int64")
		case string:
			fmt.Println(arg, "is string")
		default:
			fmt.Println(arg, "is unknown type")
		}
	}
}

func testAdd() {
	fmt.Println("请输入两个需要相加的数字: ")
	var first int
	var second int
	fmt.Scanln(&first, &second)
	sum, err := add(first, second)
	if err != nil {
		fmt.Println("相加操作异常！", err)
	} else {
		fmt.Printf("%d + %d = %d\n", first, second, sum)
	}
}

// 多返回值
func add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("不可为负数！")
		return
	}

	return a + b, nil
}
