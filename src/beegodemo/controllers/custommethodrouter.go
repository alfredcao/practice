package controllers

import "github.com/astaxie/beego"

type CustomMethodRouterController struct {
	beego.Controller
}

func (p *CustomMethodRouterController) Get() {
	p.Ctx.WriteString("invoke Get()")
}

func (p *CustomMethodRouterController) Post() {
	p.Ctx.WriteString("invoke Post()")
}

func (p *CustomMethodRouterController) CustomMethod() {
	p.Ctx.WriteString("invoke CustomMethod()")
}

func (p *CustomMethodRouterController) CustomPostMethod() {
	p.Ctx.WriteString("invoke CustomPostMethod()")
}
