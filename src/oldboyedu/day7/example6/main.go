package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src, srcOpenErr := os.Open("var/file/data.txt")
	if srcOpenErr != nil {
		fmt.Println("打开源文件失败！", srcOpenErr)
		return
	}

	defer src.Close()

	dst, dstOpenErr := os.OpenFile("var/file/dstData.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if dstOpenErr != nil {
		fmt.Println("打开目标文件失败！", dstOpenErr)
		return
	}

	defer dst.Close()

	n, copyErr := io.Copy(dst, src)

	if copyErr != nil {
		fmt.Println("拷贝文件失败！", copyErr)
		return
	} else {
		fmt.Printf("拷贝文件成功，共拷贝 %d 字节！\n", n)
	}
}
