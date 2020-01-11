package main

import (
	"fmt"
	tail2 "github.com/hpcloud/tail"
	"time"
)

func main() {
	filePath := "/Users/caozhen/dev/go/practice/src/oldboyedu/day12/example2/nginx.log"
	tail, err := tail2.TailFile(filePath, tail2.Config{
		ReOpen:    true,
		MustExist: false,
		Follow:    true,
		Poll:      true,
		Location: &tail2.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
	})

	if err != nil {
		fmt.Println("tail file error :", err)
		return
	}

	for {
		msg, ok := <-tail.Lines
		if !ok {
			fmt.Println("tail file close reopen, filePath :", filePath)
			time.Sleep(time.Second)
			continue
		}

		fmt.Println("read line :", msg.Text)
	}
}
