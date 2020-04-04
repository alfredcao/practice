package routers

import (
	"beegodemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/custom/config", &controllers.CustomConfigController{})
	beego.Router("/app/config", &controllers.AppConfigController{})
}
