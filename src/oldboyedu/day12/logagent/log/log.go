package log

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"oldboyedu/day12/logagent/conf"
)

func InitLog() (err error) {
	config := make(map[string]interface{})
	config["filename"] = conf.AppConfig.LogConf.Path
	config["level"] = convertLogLevel(conf.AppConfig.LogConf.Level)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("init log failed, json marshal failed, err:", err)
		return
	}

	err = logs.SetLogger(logs.AdapterFile, string(configStr))

	return
}

func convertLogLevel(logLevel string) int {
	switch logLevel {
	case "debug":
		return logs.LevelDebug
	case "trace":
		return logs.LevelTrace
	case "info":
		return logs.LevelInfo
	case "warn":
		return logs.LevelWarn
	case "error":
		return logs.LevelError
	}
	return logs.LevelDebug
}
