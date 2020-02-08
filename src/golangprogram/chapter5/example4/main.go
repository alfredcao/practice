package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	taskChan       = make(chan int, 10)
	taskCount      = 10
	goroutineCount = 4
	waitGroup      sync.WaitGroup
)

func main() {
	rand.Seed(time.Now().UnixNano())
	waitGroup.Add(goroutineCount)
	for i := 1; i <= goroutineCount; i++ {
		go doTask(i)
	}

	for i := 1; i <= taskCount; i++ {
		taskChan <- i
	}

	close(taskChan)
	waitGroup.Wait()
	fmt.Printf("all goroutine finish\n")
}

func doTask(id int) {
	defer waitGroup.Done()
	fmt.Printf("goroutine[%d] start\n", id)

	for {
		task, ok := <-taskChan
		if !ok {
			fmt.Printf("goroutine[%d] shutdown\n", id)
			break
		}

		fmt.Printf("goroutine[%d] do task[%d]\n", id, task)
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		fmt.Printf("goroutine[%d] finish task[%d]\n", id, task)
	}
}
