package main

import "fmt"

func main() {
	var a = 1
	fmt.Println("&a : ", &a)

	// 无论是方法调用还是赋值操作，指针类型都是值传递
	p := &a
	fmt.Println("p : ", p)
	fmt.Println("&p : ", &p)
	pointer(p)
	q := p
	fmt.Println("q : ", q)
	fmt.Println("&q : ", &q)
}

func pointer(pointer *int) {
	fmt.Println("&pointer : ", &pointer)
}
