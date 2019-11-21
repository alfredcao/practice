package main

import (
	"fmt"
	"time"
)

const (
	Man   = 1
	Woman = 2
)

func main() {
	for {
		unix := time.Now().Unix()
		if (unix % Woman) == 0 {
			fmt.Println("woman")
		} else {
			fmt.Println("man")
		}

		time.Sleep(1 * time.Second)
	}
}
