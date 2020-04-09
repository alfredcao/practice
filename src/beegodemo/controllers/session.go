package controllers

import "github.com/astaxie/beego"

type SessionController struct {
	beego.Controller
}

func (p *SessionController) Get() {
	seq := p.GetSession("seq")
	if seq == nil {
		p.SetSession("seq", 1)
	} else {
		p.SetSession("seq", seq.(int)+1)
	}
}
