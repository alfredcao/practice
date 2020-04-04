package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type CustomConfigController struct {
	beego.Controller
}

func (p *CustomConfigController) Get() {
	configKey := p.GetString("configKey")
	if configKey == "" {
		p.Ctx.WriteString("please pass valid configKey")
		return
	}
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
