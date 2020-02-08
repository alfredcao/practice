package main

import "fmt"

func main() {
	t := Time{1}
	fmt.Printf("Init Time Address : %p\n", &t)

	t2 := t.Add(1)
	fmt.Printf("Return Time Address : %p\n", &t2)

}

type Time struct {
	sec int
}

func (t Time) Add(secInterval int) Time {
	fmt.Printf("Add Time Address : %p\n", &t)
	t.sec += secInterval
	return t
}
