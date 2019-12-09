package main

import "fmt"

func main() {
	//fab(10)
	testArray()
}

func fab(n int) {
	var a = make([]int, n)
	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
}

func testArray() {
	var a1 [5]int = [5]int{1, 2, 3, 4, 5}
	var a2 = [5]int{1, 2, 3, 4, 5}
	var a3 = [...]int{12, 234, 234, 36}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}
