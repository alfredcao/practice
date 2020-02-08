package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count     int64
	waitGroup sync.WaitGroup
)

func main() {
	waitGroup.Add(2)
	runtime.GOMAXPROCS(1)
	go incr(1)
	go incr(2)

	waitGroup.Wait()
	fmt.Println("go routine finish, count :", count)
}

//func incr(id int)  {
//	defer waitGroup.Done()
//	fmt.Printf("go routine[%d] start run\n", id)
//
//	for i := 0; i < 2; i++ {
//
//		value := count
//
//		fmt.Printf("go routine[%d] yield\n", id)
//		runtime.Gosched()
//		fmt.Printf("go routine[%d] rerun\n", id)
//
//		value++
//		count = value
//
//	}
//}

func incr(id int) {
	defer waitGroup.Done()
	fmt.Printf("go routine[%d] start run\n", id)

	for i := 0; i < 2; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Printf("go routine[%d] yield\n", id)
		runtime.Gosched()
	}
}
