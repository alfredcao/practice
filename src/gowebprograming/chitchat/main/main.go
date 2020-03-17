package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/julienschmidt/httprouter"
	"gowebprograming/chitchat/controller"
	_ "gowebprograming/chitchat/dao"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/threadList", accessLog(controller.GetThreadList))
	router.GET("/thread/:id", accessLog(controller.GetThreadById))
	server := http.Server{
		Addr: "localhost:8080",
		//Handler : &HelloWorldHandler{},
		Handler: router,
	}

	//http.Handle("/hello", &HelloHandler{})
	//http.Handle("/world", &WorldHandler{})
	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/world", world)
	err := server.ListenAndServe()
	logs.Error("start server failed, err :", err)
}

type HelloWorldHandler struct{}

func (p *HelloWorldHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

type HelloHandler struct{}

func (p *HelloHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello")
}

type WorldHandler struct{}

func (p *WorldHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "World")
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello")
}

func world(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "World")
}

func accessLog(h httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		fmt.Printf("access %s\n", req.URL.Path)
		h(resp, req, params)
	})
}
