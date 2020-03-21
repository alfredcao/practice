package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/julienschmidt/httprouter"
	"gowebprograming/chitchat/controller"
	_ "gowebprograming/chitchat/dao"
	"io/ioutil"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/threadList", accessLog(controller.GetThreadList))
	router.GET("/thread/:id", accessLog(controller.GetThreadById))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	logs.Error("start server failed, err :", err)
}

func testMain() {
	router := httprouter.New()
	router.POST("/params", accessLog(params))
	router.POST("/body", accessLog(body))
	router.GET("/writeBody", accessLog(writeBody))
	router.GET("/writeHeader", accessLog(writeHeader))
	router.GET("/header", accessLog(header))
	router.GET("/json", accessLog(jsonExample))
	router.GET("/setCookie", accessLog(setCookie))
	router.GET("/getCookie", accessLog(getCookie))
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

func params(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	//req.ParseForm()
	//fmt.Fprint(resp, req.Form)
	//fmt.Fprint(resp, req.PostForm)

	//req.ParseMultipartForm(1024)
	//fmt.Fprint(resp, req.MultipartForm)

	//fmt.Fprint(resp, req.FormValue("name"))

	file, _, err := req.FormFile("file")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprint(resp, string(data))
		}
	}
}

func body(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	body := make([]byte, req.ContentLength)
	req.Body.Read(body)
	fmt.Fprint(resp, string(body))
}

func writeBody(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	resp.WriteHeader(200)
	body := `
<html>
	<head>
		<title>welcome to go</title>
	</head>
	<body>
		<p>hello, what's your name?</p>
	</body>
</html>`
	resp.Write([]byte(body))
}

func writeHeader(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	resp.WriteHeader(501)
	resp.Write([]byte("not implement, please try other api!"))
}

func header(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	resp.Header().Set("Location", "http://www.google.com")
	resp.WriteHeader(302)
}

func jsonExample(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	resp.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "caozhen",
		Threads: []string{"one", "two", "three"},
	}
	jsonStr, _ := json.Marshal(post)
	resp.Write([]byte(jsonStr))
	resp.WriteHeader(200)
}

func setCookie(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Learn How to set cookie",
		HttpOnly: true,
	}

	//resp.Header().Set("Set-Cookie", c1.String())
	//resp.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(resp, &c1)
	http.SetCookie(resp, &c2)
}

func getCookie(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	//cookie := req.Header.Get("Cookie")
	cookie := req.Cookies()
	fmt.Println(cookie)
	fmt.Println(req.Cookie("first_cookie"))
	fmt.Println(req.Cookie("not_exist_cookie"))
}

type Post struct {
	User    string   `json:"user"`
	Threads []string `json:"threads"`
}

func accessLog(h httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		fmt.Printf("access %s\n", req.URL.Path)
		h(resp, req, params)
	})
}
