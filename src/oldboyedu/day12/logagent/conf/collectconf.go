package conf

import (
	"context"
	"encoding/json"
	"fmt"
	"oldboyedu/day12/logagent/etcd"
	"oldboyedu/day12/logagent/util"
	"time"
)

const collectConfPrefixKey = "/logagent/collect/conf/"

func loadColectConf() (err error) {
	for _, localIp := range util.LocalIps {
		// 获取本机IP相关收集配置
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		getRes, getErr := etcd.EtcdCli.Get(ctx, collectConfPrefixKey+localIp)
		cancel()
		if getErr != nil {
			err = getErr
			return
		}

		for _, kv := range getRes.Kvs {
			key := string(kv.Key)
			value := string(kv.Value)
			fmt.Printf("load collect config success, key : %s, value : %s\n", key, value)
			var collectConfs = make([]AppCollectConf, 0)

			unmarshalErr := json.Unmarshal(kv.Value, &collectConfs)
			if unmarshalErr != nil {
				fmt.Println("unmarshal failed, err :", err)
				continue
			} else if len(collectConfs) != 0 {
				AppConfig.CollectConfs = append(AppConfig.CollectConfs, collectConfs...)
			}
		}
	}
	return
}
