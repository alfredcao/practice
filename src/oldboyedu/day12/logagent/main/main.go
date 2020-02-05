package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"oldboyedu/day12/logagent/conf"
	"oldboyedu/day12/logagent/es"
	"oldboyedu/day12/logagent/kafka"
	"oldboyedu/day12/logagent/log"
	"oldboyedu/day12/logagent/tailf"
	_ "oldboyedu/day12/logagent/util"
	"os"
	"time"
)

func main() {
	confType := os.Getenv("LOG_AGENT_CONF_TYPE")
	if len(confType) == 0 {
		confType = "ini"
	}
	confFilePath := os.Getenv("LOG_AGENT_CONF_FILE_PATH")
	if len(confFilePath) == 0 {
		confFilePath = "./var/conf/logagent.ini"
	}

	// 加载配置
	err := conf.LoadConf(confType, confFilePath)
	if err != nil {
		fmt.Println("load config failed, err :", err)
		return
	}

	fmt.Printf("load config success, data :%v\n", conf.AppConfig)

	// 初始化日志组件
	err = log.InitLog()
	if err != nil {
		fmt.Println("init log failed, err :", err)
		return
	}
	fmt.Println("init log success")
	logs.Info("init log success")

	// 初始化ES组件
	err = es.InitES()
	if err != nil {
		fmt.Println("init es failed, err :", err)
		return
	}

	fmt.Println("init es success")
	logs.Info("init es success")

	// 初始化tail组件
	err = tailf.InitTailf()
	if err != nil {
		fmt.Println("init tail failed, err :", err)
		return
	}

	fmt.Println("init tail success")
	logs.Info("init tail success")

	// 初始化kafka组件
	err = kafka.InitKafkaProducer()
	if err != nil {
		fmt.Println("init kafka producer failed, err :", err)
		return
	}

	fmt.Println("init kafka producer success")
	logs.Info("init kafka producer success")
	err = kafka.InitKafkaConsumer()
	if err != nil {
		fmt.Println("init kafka consumer failed, err :", err)
		return
	}

	fmt.Println("init kafka consumer success")
	logs.Info("init kafka consumer success")

	go func() {
		fd, _ := os.OpenFile("./var/log/nginx.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

		for {
			now := time.Now().Format("2006-01-02 15:04:05")
			content := "====== " + now + " =====\n"
			buf := []byte(content)
			fd.Write(buf)
			time.Sleep(time.Second)
		}

		fd.Close()
	}()

	time.Sleep(time.Hour)
}
