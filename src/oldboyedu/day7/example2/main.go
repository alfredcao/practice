package main

import "fmt"

func main() {
	str := "stu1 18 89.1"
	format := "%s %d %f"
	var stu student

	fmt.Sscanf(str, format, &stu.name, &stu.age, &stu.score)
	fmt.Println(stu)
}

type student struct {
	name  string
	age   int
	score float32
}
