package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"time"
)

func main() {
	//testGet()
	//testDebug()
	//testTimeout()
	//testPost()
	testUploadFile()
}

func testGet() {
	request := httplib.Get("http://www.baidu.com")
	request.Header("Accept-Encoding", "gzip,deflate,sdch")
	request.Header("Host", "baidu.com")
	request.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	response, err := request.Response()
	if err != nil {
		fmt.Println("http request failed, err :", err)
		return
	}
	fmt.Println(response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func testDebug() {
	httplib.Get("http://www.baidu.com").Debug(true).Response()
}

// http request failed, err : Get http://www.baidu.com: dial tcp: i/o timeout
func testTimeout() {
	req := httplib.Get("http://www.baidu.com")
	req.SetTimeout(time.Nanosecond, time.Nanosecond)
	body, err := req.String()
	if err != nil {
		fmt.Println("http request failed, err :", err)
		return
	}
	fmt.Println(body)
}

func testPost() {
	req := httplib.Post("http://www.baidu.com")
	req.Param("name", "caozhen")
	resp, err := req.Response()
	if err != nil {
		fmt.Println("http request failed, err :", err)
		return
	}

	fmt.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func testUploadFile() {
	req := httplib.Post("http://www.baidu.com")
	req.Param("name", "caozhen")
	req.PostFile("file", "test.txt")
	req.Response()
}
