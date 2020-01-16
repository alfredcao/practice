package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		time.Sleep(time.Second)
	}()

	for num := range gen(ctx) {
		fmt.Println(num)
		if num >= 5 {
			break
		}
	}
}

func gen(ctx context.Context) chan int {
	var intChan = make(chan int)
	go func() {
		var num = 0
		for true {
			select {
			case <-ctx.Done():
				fmt.Println("go route exist")
				return
			case intChan <- num:
				num++
			}
		}
	}()

	return intChan
}
