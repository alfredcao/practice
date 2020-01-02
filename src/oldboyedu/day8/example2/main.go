package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)

	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan)
	}

	go func() {
		for i := 0; i < 100000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	go func() {
		for result := range resultChan {
			fmt.Println(result)
		}
	}()
	time.Sleep(time.Second * 10)

}

func calc(intChan chan int, resultChan chan int) {
	for v := range intChan {
		if isPrime(v) {
			resultChan <- v
		}
	}
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
