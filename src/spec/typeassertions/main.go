package main

import (
	"fmt"
)

type I interface{ m() }

func main() {
	var x interface{} = 7
	v1, ok := x.(I)
	fmt.Println(ok, v1)
}
