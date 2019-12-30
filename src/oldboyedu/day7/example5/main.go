package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, oe := os.OpenFile("var/file/data.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if oe != nil {
		fmt.Println("打开文件失败！", oe)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		n, we := writer.WriteString("hello world!\n")
		if we != nil {
			fmt.Println("写入文件失败！", we)
		} else {
			fmt.Printf("成功写入 %d 字节！\n", n)
		}
	}
	fe := writer.Flush()
	if fe != nil {
		fmt.Println("内容刷入文件失败！", fe)
	} else {
		fmt.Println("内容刷入文件成功！")
	}
}
