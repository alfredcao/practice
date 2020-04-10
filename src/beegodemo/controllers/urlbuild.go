package controllers

import "github.com/astaxie/beego"

type UrlBuildController struct {
	beego.Controller
}

func (p *UrlBuildController) Get() {
	p.Data["Username"] = "astaxie"
	p.Ctx.Output.Body([]byte("ok"))
}

func (p *UrlBuildController) List() {
	p.Ctx.Output.Body([]byte("list method"))
}

func (p *UrlBuildController) Params() {
	p.Ctx.Output.Body([]byte(p.Ctx.Input.Params()["0"] + p.Ctx.Input.Params()["1"] + p.Ctx.Input.Params()["2"]))
}

func (p *UrlBuildController) Myext() {
	p.Ctx.Output.Body([]byte(p.Ctx.Input.Param(":ext")))
}

func (p *UrlBuildController) GetUrl() {
	p.Ctx.Output.Body([]byte(p.URLFor(".Myext")))
}
