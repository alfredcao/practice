package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// 求100以内的所有质数
	fmt.Println("以下为100以内的质数：")
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Printf("%d\n", i)
		}
	}

	// 求 100 - 999 之间的水仙花数
	fmt.Println("以下为 100 - 999 的水仙花数：")
	for i := 100; i <= 999; i++ {
		if isDaffodil(i) {
			fmt.Printf("%d\n", i)
		}
	}

	factorialSum := factorialSum(4)
	fmt.Printf("10以内所有数字的阶乘和是：%d", factorialSum)
}

// 判断数字是否为质数
func isPrime(num int) bool {
	if num < 2 || (num%2 == 0) {
		return false
	}

	top := int(math.Sqrt(float64(num)))

	for i := 2; i <= top; i++ {
		if (num % i) == 0 {
			return false
		}
	}

	return true
}

// 判断数字是否为水仙花数
func isDaffodil(num int) bool {
	numStr := fmt.Sprintf("%d", num)
	numLen := len(numStr)

	total := 0
	_3 := float64(3)
	for i := 0; i < numLen; i++ {
		tmp, _ := strconv.ParseInt(numStr[i:i+1], 10, 32)
		total += int(math.Pow(float64(tmp), _3))
	}

	if total == num {
		return true
	}

	return false
}

// 获取num之前所有数字的阶乘和
func factorialSum(num int) int {
	total := 0

	for i := 1; i <= num; i++ {
		total += factorial(i)
	}

	return total
}

func factorial(num int) int {
	total := 0
	if num == 1 {
		total = 1
	} else {
		total += num * factorial(num-1)
	}

	return total
}
