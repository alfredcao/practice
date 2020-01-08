package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("0.0.0.0:9999", nil)
	//err := http.ListenAndServe("127.0.0.1:9999", nil)
	//err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		fmt.Println("HTTP服务器启动失败！", err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "hello world!")
}
