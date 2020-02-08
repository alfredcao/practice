package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown  int64
	waitGroup sync.WaitGroup
)

func main() {
	waitGroup.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(time.Second * 3)
	atomic.StoreInt64(&shutdown, 1)
	waitGroup.Wait()
}

func doWork(name string) {
	defer waitGroup.Done()
	for {
		fmt.Printf("do work %s\n", name)
		time.Sleep(500 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("work %s finish\n", name)
			break
		}
	}
}
