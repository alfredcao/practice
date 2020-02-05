package conf

import (
	"errors"
	"github.com/astaxie/beego/config"
	"oldboyedu/day12/logagent/etcd"
)

var (
	AppConfig = AppConf{}
)

type AppConf struct {
	LogConf      AppLogConf       `json:"logConf"`
	CollectConfs []AppCollectConf `json:"collectConfs"`
	KafkaConf    AppKafkaConf     `json:"kafkaConf"`
	EtcdConf     AppEtcdConf      `json:"etcdConf"`
	ESConf       AppESConf        `json:"esConf"`
}

type AppLogConf struct {
	Path  string `json:"path"`
	Level string `json:"level"`
}

type AppCollectConf struct {
	LogPath  string `json:"logPath"`
	LogTopic string `json:"logTopic"`
}

type AppKafkaConf struct {
	Addr string `json:"addr"`
}

type AppEtcdConf struct {
	Addr string `json:"addr"`
}

type AppESConf struct {
	Addr string `json:"addr"`
}

func LoadConf(confType, confFilePath string) (err error) {
	// 初始化配置器
	configer, err := config.NewConfig(confType, confFilePath)
	if err != nil {
		return
	}

	// 加载日志相关配置
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

	// 加载ETCD相关配置
	etcdAddr := configer.String("etcd::addr")
	if len(etcdAddr) == 0 {
		err = errors.New("no etcd addr")
		return
	}

	AppConfig.EtcdConf = AppEtcdConf{
		Addr: etcdAddr,
	}

	err = etcd.InitEtcd([]string{AppConfig.EtcdConf.Addr})
	if err != nil {
		return
	}

	// 加载kafka相关配置
	kafkaAddr := configer.String("kafka::addr")
	if len(kafkaAddr) == 0 {
		err = errors.New("no kafka address")
		return
	}
	AppConfig.KafkaConf = AppKafkaConf{
		Addr: kafkaAddr,
	}

	// 加载es相关配置
	esAddr := configer.String("es::addr")
	if len(esAddr) == 0 {
		err = errors.New("no es address")
		return
	}
	AppConfig.ESConf = AppESConf{
		Addr: esAddr,
	}

	// 加载收集信息相关配置
	//collectLogTopic := configer.String("collect::log_topic")
	//if len(collectLogTopic) == 0 {
	//	collectLogTopic = "nginx_log"
	//}
	//collectLogPath := configer.String("collect::log_path")
	//if len(collectLogPath) == 0 {
	//	err = errors.New("no collect log path")
	//	return
	//}
	//AppConfig.CollectConfs = append(AppConfig.CollectConfs, AppCollectConf{
	//	LogPath:  collectLogPath,
	//	LogTopic: collectLogTopic,
	//})

	loadColectConf()

	return
}
