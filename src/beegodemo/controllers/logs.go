package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LogsController struct {
	beego.Controller
}

func (p *LogsController) Get() {
	logs.Debug("debug log")
	logs.Info("info log")
	logs.Notice("notice log")
	logs.Warn("warn log")
	logs.Error("error log")
	logs.Critical("critical log")
	logs.Alert("alert log")
	logs.Emergency("emergency log")
}
