package es

import (
	"github.com/astaxie/beego/logs"
	"github.com/olivere/elastic"
	"oldboyedu/day12/logagent/conf"
)

var ESClient *elastic.Client

func InitES() (err error) {
	cli, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(conf.AppConfig.ESConf.Addr))
	if err != nil {
		logs.Error("new elasticsearch client failed, err :", err)
		return
	}

	ESClient = cli
	return
}
