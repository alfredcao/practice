package main

import (
	"fmt"
	"html/template"
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
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprint(res, "打开模板文件失败！", err)
	} else {
		t.Execute(res, Person{
			Title: "我的个人网站",
			Name:  "Lilei",
			age:   18,
		})
	}

}

type Person struct {
	Title string
	Name  string
	age   int
}
