package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type ErrorHandleController struct {
	beego.Controller
}

func (p *ErrorHandleController) AbortDemo() {
	// beego 框架默认支持 401、403、404、500、503 这几种错误的处理
	p.Abort("401")

	p.Ctx.WriteString("no abort")
}

func (p *ErrorHandleController) DBError() {
	p.Abort("DBError")

	p.Ctx.WriteString("no abort")
}

func (p *ErrorHandleController) RedisError() {
	p.Abort("RedisError")

	p.Ctx.WriteString("no abort")
}

func (p *ErrorHandleController) ErrorRedisError() {
	p.Ctx.WriteString("redis error, please contact admin!")
}

func Custom404ErrorHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("404 not found, please check url!"))
}

func DBErrorHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("db error, please contact admin!"))
}
