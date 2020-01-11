package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "/Users/caozhen/dev/go/practice/src/oldboyedu/day12/example4/nginx.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.Debug("this is a debug log!")
	logs.Trace("this is a trace log!")
	logs.Info("this is a info log!")

}
