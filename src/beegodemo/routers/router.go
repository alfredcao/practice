package routers

import (
	"beegodemo/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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

}

type BasicRouterHandler struct{}

func (p *BasicRouterHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("success"))
}
