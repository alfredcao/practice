package controllers

import "github.com/astaxie/beego"

type RegexpRouterController struct {
	beego.Controller
}

func (p *RegexpRouterController) Get() {
	p.Ctx.WriteString(p.Ctx.Input.Param(":id"))
}
