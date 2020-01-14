package conf

import (
	"errors"
	"github.com/astaxie/beego/config"
)

var (
	AppConfig = AppConf{}
)

type AppConf struct {
	LogConf      AppLogConf
	CollectConfs []AppCollectConf
	KafkaConf    AppKafkaConf
}

type AppLogConf struct {
	Path  string
	Level string
}

type AppCollectConf struct {
	LogPath  string
	LogTopic string
}

type AppKafkaConf struct {
	Addr string
}

func LoadConf(confType, confFilePath string) (err error) {
	configer, err := config.NewConfig(confType, confFilePath)
	if err != nil {
		return
	}

	logPath := configer.String("log::path")
	if len(logPath) == 0 {
		logPath = "./var/log/logagent.log"
	}

	logLevel := configer.String("log::level")
	if len(logLevel) == 0 {
		logLevel = "debug"
	}

	AppConfig.LogConf = AppLogConf{
		Path:  logPath,
		Level: logLevel,
	}

	collectLogPath := configer.String("collect::log_path")
	if len(collectLogPath) == 0 {
		err = errors.New("no collect log path")
		return
	}

	collectLogTopic := configer.String("collect::log_topic")
	if len(collectLogTopic) == 0 {
		collectLogTopic = "nginx_log"
	}
	AppConfig.CollectConfs = append(AppConfig.CollectConfs, AppCollectConf{
		LogPath:  collectLogPath,
		LogTopic: collectLogTopic,
	})

	kafkaAddr := configer.String("kafka::addr")
	if len(kafkaAddr) == 0 {
		err = errors.New("no kafka address")
		return
	}
	AppConfig.KafkaConf = AppKafkaConf{
		Addr: kafkaAddr,
	}

	return
}
