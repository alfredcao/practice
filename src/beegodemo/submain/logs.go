package main

import "github.com/astaxie/beego/logs"

func main() {
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","maxlines":2,"maxsize":1024000,"daily":true,"maxdays":7,"rotate":true,"level":7,"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	// 异步输出日志
	//logs.Async(1e3)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	logs.Emergency("test emergency log")
	logs.Alert("test alert log")
	logs.Warn("test warn log")
	logs.Info("test info log")
	logs.Debug("test debug log")
}
