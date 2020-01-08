package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", logPanic(hello))
	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		fmt.Println("HTTP服务器启动失败！", err)
	}
}

func logPanic(handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("发生异常！", err)
			}
		}()

		handle(writer, request)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "hello world!")
	panic("some thing error")
}
