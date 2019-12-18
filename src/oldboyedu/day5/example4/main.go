package main

import "fmt"

func main() {
	stu1 := Student{}
	stu1.Name = "stu1"
	stu1.Age = 18
	stu1.Class = "class 1"
	stu1.Score = 80

	fmt.Println(stu1)
}

type People struct {
	Name string
	Age  int
}

type Student struct {
	People
	Class string
	Score int
}
