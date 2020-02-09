package main

import (
	"fmt"
	"golangprogram/chapter7/runner"
	"os"
	"time"
)

var timeout = time.Second * 3

func main() {
	r := runner.NewRunner(timeout)
	r.Add(CreatTask())
	//r.Add(CreatTask())
	//r.Add(CreatTask())
	err := r.Start()

	if err != nil {
		switch err {
		case runner.ErrInterrupt:
			fmt.Println("tasks finish by interrupt")
			os.Exit(1)
		case runner.ErrTimeout:
			fmt.Println("tasks finish by timeout")
			os.Exit(2)
		}
	}

	fmt.Println("tasks finish by normal")
}

func CreatTask() func(int) {
	return func(id int) {
		fmt.Printf("task[%d] start\n", id)
		time.Sleep(time.Duration(id) * time.Second)
		fmt.Printf("task[%d] finish\n", id)
	}
}
