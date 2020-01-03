package main

import (
	"fmt"
	"math"
)

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	go func() {
		for i := 0; i < 100000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}

}

func calc(intChan chan int, resultChan chan int, exitChan chan bool) {
	for v := range intChan {
		if isPrime(v) {
			resultChan <- v
		}
	}

	exitChan <- true
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
