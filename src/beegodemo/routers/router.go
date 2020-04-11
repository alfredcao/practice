package routers

import (
	"beegodemo/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"net/http"
)

func init() {
	// 基础路由
	beego.Get("/basicrouter/get", func(context *context.Context) {
		context.WriteString("success")
	})
	beego.Post("/basicrouter/post", func(context *context.Context) {
		context.WriteString("success")
	})
	beego.Any("/basicrouter/any", func(context *context.Context) {
		context.WriteString("success")
	})
	beego.Handler("/basicrouter/handler", &BasicRouterHandler{})

	// 固定路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/custom/config", &controllers.CustomConfigController{})
	beego.Router("/app/config", &controllers.AppConfigController{})

	// 正则路由
	beego.Router("/regexprouter/?:id", &controllers.RegexpRouterController{})
	//beego.Router("/regexprouter/:id", &controllers.RegexpRouterController{})
	//beego.Router("/regexprouter/:id([0-9]+)", &controllers.RegexpRouterController{})
	//beego.Router("/regexprouter/:id(\\w+)", &controllers.RegexpRouterController{})

	// 自定义方法及 RESTful 规则
	beego.Router("/custommethodrouter", &controllers.CustomMethodRouterController{}, "get,delete:CustomMethod;post:CustomPostMethod")

	// 自动匹配
	beego.AutoRouter(&controllers.AutoRouterController{})

	// 获取参数
	beego.Router("/param", &controllers.ParamController{})

	// session控制
	beego.Router("/session", &controllers.SessionController{})

	// 过滤器
	beego.Router("/login", &controllers.LoginController{})
	//beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
	//	fmt.Println("请求URI :", ctx.Request.RequestURI)
	//	if ctx.Request.URL.Path != "/login" {
	//		_, ok := ctx.Input.Session("uid").(int)
	//		if !ok {
	//			ctx.Redirect(http.StatusFound, "/login")
	//		}
	//	}
	//})

	// URL配置
	beego.Router("/urlbuild/list", &controllers.UrlBuildController{}, "*:List")
	beego.Router("/urlbuild/:lastname/:firstname", &controllers.UrlBuildController{})
	beego.AutoRouter(&controllers.UrlBuildController{})
	fmt.Println(beego.URLFor("UrlBuildController.List"))
	fmt.Println(beego.URLFor("UrlBuildController.Get"), ":lastname", "zhen", ":firstname", "cao")
	fmt.Println(beego.URLFor("UrlBuildController.Myext"))
	fmt.Println(beego.URLFor("UrlBuildController.GetUrl"))

	// 多种数据格式输出
	beego.Router("/formatoutput/json", &controllers.FormatOutputController{}, "get:Json")
	beego.Router("/formatoutput/xml", &controllers.FormatOutputController{}, "get:Xml")
	beego.Router("/formatoutput/jsonp", &controllers.FormatOutputController{}, "get:Jsonp")

	// 表单数据验证
	beego.Router("/formvalid/valid", &controllers.FormValidController{}, "post:Valid")
	beego.Router("/formvalid/validbystructtag", &controllers.FormValidController{}, "post:ValidByStructTag")

	// 错误处理
	beego.Router("/errorhandle/abortdemo", &controllers.ErrorHandleController{}, "*:AbortDemo")
	beego.Router("/errorhandle/dberror", &controllers.ErrorHandleController{}, "*:DBError")
	beego.Router("/errorhandle/rediserror", &controllers.ErrorHandleController{}, "*:RedisError")
	beego.ErrorHandler("404", controllers.Custom404ErrorHandler)
	beego.ErrorHandler("DBError", controllers.DBErrorHandler)
	beego.ErrorController(&controllers.ErrorHandleController{})

	// 日志处理
	logs.SetLevel(logs.LevelInfo)
	beego.Router("/logs", &controllers.LogsController{})
}

type BasicRouterHandler struct{}

func (p *BasicRouterHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("success"))
}
