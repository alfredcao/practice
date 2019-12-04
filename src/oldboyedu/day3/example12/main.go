package main

import "fmt"

type op_func func(int, int) int

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	c := add
	fmt.Println(c)

	sum := op(c, 1, 2)
	fmt.Println(sum)

	d := add
	fmt.Println(d)
}

func op(op op_func, num1 int, num2 int) int {
	fmt.Println(op)
	return op(num1, num2)
}
