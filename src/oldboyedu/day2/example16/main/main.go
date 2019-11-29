package main

import "fmt"

func main() {
	var a = 100
	var b = false
	var c byte = 'x'
	var f float32 = 234.24
	fmt.Printf("a=%v\n", a)
	fmt.Printf("a=%b\n", a) // 二进制
	fmt.Printf("a=%o\n", a) // 八进制
	fmt.Printf("a=%x\n", a) // 十六进制
	fmt.Printf("a Type is %T\n", a)
	fmt.Printf("b=%v\n", b)
	fmt.Printf("b=%t\n", b)
	fmt.Printf("b Type is %T\n", b)
	fmt.Printf("c=%v\n", c)
	fmt.Printf("c=%c\n", c)
	fmt.Printf("c Type is %T\n", c)
	fmt.Printf("f=%f\n", f)
}
