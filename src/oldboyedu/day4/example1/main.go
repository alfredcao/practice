package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	i := new(int) // new返回指针类型
	fmt.Printf("new(int) : Type = %T, value = %v\n", i, i)

	s1 := make([]int, 3, 10)
	fmt.Printf("make([]int, 3, 10) : Type = %T, value = %v\n", s1, s1)
	s1 = append(s1, -1)
	s2 := []int{1, 10, 99}
	s1 = append(s1, s2...)

	fmt.Println("append ", s1)

	s3 := new([]int)
	fmt.Printf("new([]int) : Type = %T, value = %v\n", s3, s3)
	(*s3)[0] = 1

	fmt.Println("s3 ", s3)
}
