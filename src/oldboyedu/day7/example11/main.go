package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := OpenFile("/tmp/sfsf.txt")
	if err != nil {
		fmt.Println("打开文件错误！", err)
		return
	}

	defer file.Close()

	fmt.Println("打开文件成功！")
}

func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		err = &PathError{
			path:       path,
			op:         "read",
			message:    err.Error(),
			createTime: time.Now().String(),
		}
	}

	return file, err
}

type PathError struct {
	path       string
	op         string
	message    string
	createTime string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path = %s, op = %s, message = %s, createTime = %s", p.path, p.op, p.message, p.createTime)
}
