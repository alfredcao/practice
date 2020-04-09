package main

import (
	_ "beegodemo/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
)

func main() {
	beego.Run()
}
