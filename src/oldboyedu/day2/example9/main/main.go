package main

import "fmt"

var a = "G"

func main() {
	first := 100
	second := 200

	//firstAdd := &first
	//secondAdd := &second
	//swag(firstAdd, secondAdd)

	//first, second = swag1(first, second)
	first, second = second, first

	fmt.Println("first=", first)
	fmt.Println("second=", second)

	n()
	m()
}

func swag(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func swag1(a int, b int) (int, int) {
	return b, a
}

func n() {
	fmt.Println("a=", a)
}

func m() {
	a := "O"
	fmt.Println("a=", a)
	n()
}
