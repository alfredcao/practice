package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func main() {
	for _, url := range urls {
		res, err := http.Head(url)
		if err != nil {
			fmt.Printf("访问网站[%s]失败！err : %v\n", url, err)
		} else {
			fmt.Printf("访问网站[%s]成功，响应码 : %d！\n", url, res.StatusCode)
		}
	}
}
