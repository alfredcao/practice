package main

import (
	"fmt"
)

func Add(x int, y int, ch chan int) {
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
	ch <- 1
}

/**
channel变量声明语法：var chanName chan ElementType
定义一个channel：chanName = make(chan int)
*/
func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Add(i, i, chs[i])
	}

	for _, ch := range chs {
		<-ch
	}

	//time.Sleep(time.Second)
}
