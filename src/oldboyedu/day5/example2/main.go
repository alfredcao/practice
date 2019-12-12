package main

import "fmt"

func main() {
	testOneWayLink()
}

func testOneWayLink() {

	head := &Student{
		Name:  "head",
		Age:   18,
		Score: 99,
	}
	stu1 := &Student{
		Name:  "stu1",
		Age:   20,
		Score: 88,
	}
	stu2 := &Student{
		Name:  "stu2",
		Age:   19,
		Score: 95,
	}

	head.Next = stu1
	stu1.Next = stu2

	p := head
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}

type Student struct {
	Name  string
	Age   int
	Score float32
	Next  *Student
}
