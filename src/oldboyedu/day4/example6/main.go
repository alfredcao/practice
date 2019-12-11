package main

import (
	"fmt"
	"sort"
)

func main() {
	//testSlice()
	//testString()
	testSliceSort()
}

func testSliceSort() {
	var arr = []int{24, 541, 346, 6, 324}
	fmt.Println("排序前，24的索引为 : ", sort.SearchInts(arr, 24))
	sort.Ints(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("排序后，24的索引为 : ", sort.SearchInts(arr, 24))
}

func testSlice() {
	var arr = [5]int{1, 2, 3, 4, 5}
	slice := arr[1:5]
	fmt.Printf("arr : %v, slice : %v\n", arr, slice)
	fmt.Printf("arr[1] address : %p, slice address : %p\n", &arr[1], slice)
	fmt.Printf("arr len : %d, cap : %d\n", len(arr), cap(arr))
	fmt.Printf("slice len : %d, cap : %d\n", len(slice), cap(slice))

	slice = append(slice, 10)
	fmt.Printf("slice len : %d, cap : %d\n", len(slice), cap(slice))
	fmt.Printf("slice address : %p\n", slice)

	slice = append(slice, 10)
	fmt.Printf("slice len : %d, cap : %d\n", len(slice), cap(slice))
	fmt.Printf("slice address : %p\n", slice)

	slice = append(slice, 10)
	fmt.Printf("slice len : %d, cap : %d\n", len(slice), cap(slice))
	fmt.Printf("slice address : %p\n", slice)

	slice = append(slice, 10)
	fmt.Printf("slice len : %d, cap : %d\n", len(slice), cap(slice))
	fmt.Printf("slice address : %p\n", slice)

	slice = append(slice, 10)
	fmt.Printf("slice len : %d, cap : %d\n", len(slice), cap(slice))
	fmt.Printf("slice address : %p\n", slice)

	slice = append(slice, 10)
	fmt.Printf("slice len : %d, cap : %d\n", len(slice), cap(slice))
	fmt.Printf("slice address : %p\n", slice)
}

func testString() {
	str1 := "我爱你中国！"
	str2 := str1[2:3]
	fmt.Println(str1)
	fmt.Println(str2)

	r := []rune(str1)
	r[1] = '哈'

}
