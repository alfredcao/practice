package main

import "fmt"

func main() {
	var stu1 = new(Student)
	fmt.Printf("stu1: %v, %T, %p\n", stu1, stu1, stu1)
	var stu2 = Student{
		Name:  "stu2",
		Age:   10,
		Score: 89.5,
	}
	fmt.Printf("stu2: %v, %T, %p\n", stu2, stu2, &stu2)
	var stu3 = &Student{}
	stu3.Name = "stu3"
	stu3.Age = 20
	fmt.Printf("stu3: %v, %T, %p\n", stu3, stu3, stu3)
}

type Student struct {
	Name  string
	Age   int
	Score float32
}
