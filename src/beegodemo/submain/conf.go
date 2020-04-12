package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	confFile, err := config.NewConfig("ini", "src/beegodemo/submain/config.conf")
	if err != nil {
		fmt.Println("new config failed, err :", err)
		return
	}

	fmt.Println(confFile.String("AppName"))
	fmt.Println(confFile.String("demo::key1"))
	fmt.Println(confFile.String("envKey"))
}
