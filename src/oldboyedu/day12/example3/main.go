package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	configer, err := config.NewConfig("ini", "/Users/caozhen/dev/go/practice/src/oldboyedu/day12/example3/logagent.ini")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	port, err := configer.Int("server::port")
	if err != nil {
		fmt.Println("read port failed, err:", err)
		return
	} else {
		fmt.Println("port:", port)
	}

	level := configer.String("logs::level")
	fmt.Println("logs level:", level)

	path := configer.String("logs::path")
	fmt.Println("logs path:", path)

}
