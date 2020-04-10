package controllers

import "github.com/astaxie/beego"

type FormatOutputController struct {
	beego.Controller
}

type Output struct {
	Id    int
	Name  string
	Email string
}

func (p *FormatOutputController) Json() {
	output := &Output{
		Id:    123,
		Name:  "caozhen",
		Email: "caozhen@qq.com",
	}

	p.Data["json"] = output
	p.ServeJSON()
}

func (p *FormatOutputController) Xml() {
	output := &Output{
		Id:    123,
		Name:  "caozhen",
		Email: "caozhen@qq.com",
	}

	p.Data["xml"] = output
	p.ServeXML()
}

func (p *FormatOutputController) Jsonp() {
	output := &Output{
		Id:    123,
		Name:  "caozhen",
		Email: "caozhen@qq.com",
	}

	p.Data["jsonp"] = output
	p.ServeJSONP()
}
