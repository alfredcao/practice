package main

import (
	"fmt"
	"math"
)

/**
GO类型：
	- 布尔类型:bool
	- 整型:int8、byte、int16、int、uint、uintptr等
	- 浮点类型:float32、float64
	- 复数类型:complex64、complex128
	- 字符串:string
	- 字符类型:rune
	- 错误类型:error
此外，Go语言也支持以下这些复合类型:
	- 指针(pointer)
	- 数组(array)
	- 切片(slice)
	- 字典(map)
	- 通道(chan)
	- 结构体(struct)
	- 接口(interface)
*/
func main() {
	// 变量申明
	//var v1 int
	//var v2 string
	//var v3 [10]int
	//var v4 [10]int
	//var v5 struct {
	//	f int
	//}
	//var v6 *int
	//var v7 map[string] int
	//var v8 func(a int ) int
	//
	//var (
	//	v9 int
	//	v10 string
	//)

	// 变量赋值
	var a int = 0
	var i = 1
	j := 2

	a, i, j = j, i, a

	fmt.Println("a : ", a)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)

	// 匿名变量
	_, _, nickName := getName()
	fmt.Println("nickName : ", nickName)

	// 常量
	//const (
	//	pi := 3.14159265358979323846
	//	yes := true
	//	name := "caozhen"
	//)
	//pi = 1111 // Connot assign to constant

	// iota
	const (
		c0 = iota
		c1
		c2
	)

	const d0 = iota
	const e0 = iota

	fmt.Println("c0 : ", c0)
	fmt.Println("c1 : ", c1)
	fmt.Println("c2 : ", c2)
	fmt.Println("d0 : ", d0)
	fmt.Println("e0 : ", e0)

	// 星期枚举值
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	var f1 float32 = 0.0000123
	var f2 float32 = 0.000012234
	var gap float32 = 0.0000001

	fmt.Println("isEqual(0.0000123, 0.000012234, 0.0000001) : ", isEqual(f1, f2, gap))

	var cmp complex64
	//cmp = 3.2 + 1.2i
	cmp = complex(3.2, 1.2)
	fmt.Println("实部 : ", real(cmp))
	fmt.Println("虚部 : ", imag(cmp))

	var str string = "Hello World"
	fmt.Printf("字符串%s长度为 : %d\n", str, len(str))
	fmt.Printf("字符串%s首字符为 : %c\n", str, str[0])
	//str[0] = 'X' // cannot assign to str[0]

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for i, ch := range str {
		fmt.Println(i, ch)
	}

	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArray(arr)
	fmt.Println("In main(), array values:", arr)

	var myArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice = myArray[3:5]
	fmt.Println("myArray : ")
	for _, e := range myArray {
		fmt.Print(e, " ")
	}
	fmt.Println("")
	fmt.Println("mySlice : ")
	for _, e := range mySlice {
		fmt.Print(e, " ")
	}

	mySlice1 := make([]int, 5)
	mySlice2 := make([]int, 5, 10)
	fmt.Println("")
	fmt.Println("len(mySlice1)", len(mySlice1), "cap(mySlice1)", cap(mySlice1))
	fmt.Println("len(mySlice2)", len(mySlice2), "cap(mySlice2)", cap(mySlice2))

	mySlice3 := []int{1, 2, 3}
	mySlice4 := []int{4, 5, 6, 7, 8}
	copy(mySlice4, mySlice3)
	copy(mySlice3, mySlice4)
	fmt.Println("mySlice3 : ", mySlice3)
	fmt.Println("mySlice4 : ", mySlice4)

	// 数组为值类型，直接赋值或传参都会产生一次复制操作
	sourceArr := [3]int{1, 2, 3}
	targetArr := sourceArr
	sourceArr[0] = 5
	var sourceArrAddress = &sourceArr
	var targetArrAddress = &targetArr
	fmt.Println("sourceArr : ", sourceArr, ", sourceArr address : ", sourceArrAddress)
	fmt.Println("targetArr : ", targetArr, ", targetArr address : ", targetArrAddress)

	(*sourceArrAddress)[2] = 10
	fmt.Println("sourceArr : ", sourceArr)

}

func getName() (firstName, lastName, nickName string) {
	return "zhen", "cao", "alfred"
}

func isEqual(f1, f2, gap float32) bool {
	return math.Dim(float64(f1), float64(f2)) < float64(gap)
}

func modifyArray(arr [5]int) {
	arr[0] = 10
	fmt.Println("In modify(), array values:", arr)
}
