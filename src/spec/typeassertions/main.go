package main

import (
	"fmt"
	"time"
)

type I interface{ m() }

func main() {
	var x interface{} = 7
	v1, ok := x.(I)
	fmt.Println(ok, v1)

	var d = time.Second
	x = d
	v2, ok := x.(time.Duration)
	fmt.Println(ok, v2)

	switch t := x.(type) {
	case int:
		fmt.Println("int")
	case time.Duration:
		fmt.Println("time.Duration", t)
	}

}
