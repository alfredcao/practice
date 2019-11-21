package channel

import "fmt"

func TestPipe1() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	v1 := <-pipe
	fmt.Println("pipe first value : ", v1)

	pipe <- 4

	fmt.Println("pipe length : ", len(pipe))
}
