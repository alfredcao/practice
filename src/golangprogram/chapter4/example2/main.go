package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go showChannelValue(ch)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		time.Sleep(time.Second)
	}
}

func showChannelValue(ch chan int) {
	for {
		i := <-ch
		fmt.Println("从管道读取到", i)
	}
}
