package main

import (
	"fmt"
	"strings"
)

// 闭包：函数与其引用环境所组成的实体
func main() {
	add := Adder()
	fmt.Println(add(1))
	fmt.Println(add(20))
	fmt.Println(add(100))

	jpgSuffixFunc := makeSuffixFuc(".jpg")
	bmpSuffixFunc := makeSuffixFuc(".bmp")
	fmt.Println(jpgSuffixFunc("pic"))
	fmt.Println(jpgSuffixFunc("test"))
	fmt.Println(bmpSuffixFunc("pic"))
	fmt.Println(bmpSuffixFunc("test"))
}

func Adder() func(num int) int {
	var x int
	return func(num int) int {
		x += num
		return x
	}
}

func makeSuffixFuc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}
