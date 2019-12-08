package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 5)
		timeout <- true
	}()
	ch := make(chan int, 1024)

	go func() {
		rand.Seed(time.Now().UnixNano())
		r := time.Duration(rand.Intn(10)) * time.Second
		fmt.Println("随机数：", r)
		time.Sleep(r)
		ch <- 1
	}()
	select {
	case <-ch:
		fmt.Println("没超时！")
	case <-timeout:
		fmt.Println("超时啦！")
	}
}
