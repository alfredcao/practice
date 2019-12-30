package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("共有 %d 个参数！\n", len(args))
	for i, v := range args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}
}
