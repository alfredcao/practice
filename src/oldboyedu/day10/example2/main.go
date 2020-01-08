package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var urls = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func main() {
	for _, url := range urls {
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, time.Second*2)
				},
			},
		}

		//res, err := http.Head(url)
		res, err := c.Head(url)
		if err != nil {
			fmt.Printf("访问网站[%s]失败！err : %v\n", url, err)
		} else {
			fmt.Printf("访问网站[%s]成功，响应码 : %d！\n", url, res.StatusCode)
		}
	}
}
