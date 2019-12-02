package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	start := time.Now().UnixNano()
	sleep()
	end := time.Now().UnixNano()

	fmt.Printf("cost : %d\n", (end-start)/1000)

}

func sleep() {
	time.Sleep(time.Millisecond * 101)
}
