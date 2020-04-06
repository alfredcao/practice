package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type CustomConfigController struct {
	beego.Controller
}

// 每次请求都会执行
func (p *CustomConfigController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	p.Controller.Init(ctx, controllerName, actionName, app)
	fmt.Println("invoke Init()")
}

// 每次请求都会执行
func (p *CustomConfigController) Prepare() {
	fmt.Println("invoke Prepare()")
	configKey := p.GetString("configKey")
	if configKey == "" {
		p.Data["json"] = map[string]interface{}{"success": false, "msg": "please pass valid configKey"}
		p.ServeJSON()
		p.StopRun()
	}
}

// 每次请求都会执行
func (p *CustomConfigController) Finish() {
	fmt.Println("invoke Finish()")
}

func (p *CustomConfigController) Get() {
	configKey := p.GetString("configKey")
	v, err := beego.GetConfig("String", configKey, "")
	if err != nil {
		p.Ctx.WriteString(fmt.Sprintf("get config value failed, err : %v", err))
		return
	}
	p.Ctx.WriteString(v.(string))
}

type AppConfigController struct {
	beego.Controller
}

func (p *AppConfigController) Get() {
	configKey := p.GetString("configKey")
	if configKey == "" {
		p.Ctx.WriteString("please pass valid configKey")
		return
	}
	v := beego.AppConfig.String(configKey)
	p.Ctx.WriteString(v)

}
