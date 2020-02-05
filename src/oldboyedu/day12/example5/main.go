package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"oldboyedu/day12/logagent/conf"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: time.Second * 3,
	})

	if err != nil {
		fmt.Println("new etcd client failed, err :", err)
		return
	}

	var collectConfs = make([]conf.AppCollectConf, 0)
	var collectConf1 = conf.AppCollectConf{
		LogPath:  "./var/log/nginx.log",
		LogTopic: "nginx_log",
	}
	collectConfs = append(collectConfs, collectConf1)
	var collectConf2 = conf.AppCollectConf{
		LogPath:  "./var/log/error.log",
		LogTopic: "error_log",
	}
	collectConfs = append(collectConfs, collectConf2)
	byteDate, err := json.Marshal(collectConfs)
	if err != nil {
		fmt.Println("json marshal failed, err :", err)
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	putResponse, err := cli.Put(ctx, "/logagent/collect/conf/10.16.0.37", string(byteDate))
	cancel()

	if err != nil {
		fmt.Println("etcd put failed, err :", err)
		return
	} else {
		fmt.Println("etcd put success, response :", *putResponse)
	}
}
