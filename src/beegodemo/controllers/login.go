package controllers

import (
	"github.com/astaxie/beego"
	"math/rand"
)

type LoginController struct {
	beego.Controller
}

func (p *LoginController) Get() {
	p.Ctx.WriteString("登录界面")
}

func (p *LoginController) Post() {
	username := p.Input().Get("username")
	password := p.Input().Get("password")
	if username != "" && password != "" {
		p.SetSession("uid", rand.Intn(1000000))
	}
}
