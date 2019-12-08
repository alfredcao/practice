package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1024)
	chIn := (<-chan int)(ch)  // 只读channel
	chOut := (chan<- int)(ch) // 只写channel

	go func(ch <-chan int) {
		for value := range ch {
			fmt.Println("只读goroutine读取到值：", value)
		}
	}(chIn)

	go func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}(chOut)

	time.Sleep(time.Second * 15)
}
