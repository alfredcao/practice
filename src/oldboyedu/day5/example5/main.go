package main

import "fmt"

func main() {
	stu1 := Student{
		Name: "stu",
	}

	fmt.Printf("name is %s\n", stu1.getName())
}

type Student struct {
	Name string
}

func (this *Student) getName() string {
	return this.Name
}
