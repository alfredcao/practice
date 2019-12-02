package main

import "fmt"

type Rect struct {
	width, height float64
}

func (rect Rect) Area() float64 {
	return rect.width * rect.height
}

func main() {
	rect := Rect{30, 50}
	fmt.Println("rect area =", rect.Area())
}
