package controllers

import "github.com/astaxie/beego"

type AutoRouterController struct {
	beego.Controller
}

func (p *AutoRouterController) Login() {
	p.Ctx.WriteString("login")
}

func (p *AutoRouterController) Ext() {
	p.Ctx.WriteString(p.Ctx.Input.Param(":ext"))
}
