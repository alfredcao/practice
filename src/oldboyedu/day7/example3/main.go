package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/tmp/1.txt")
	if err != nil {
		fmt.Println("打开文件错误！", err)
		return
	}
	defer file.Close()

	//reader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')

	if err != nil && err != io.EOF {
		fmt.Println("读取错误！", err)
		return
	}

	fmt.Println("读取到字符串 :", str)
}
