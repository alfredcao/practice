package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"time"
)

var EtcdCli *clientv3.Client

func InitEtcd(addrs []string) (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: time.Second * 3,
	})

	if err != nil {
		return
	}

	EtcdCli = cli

	return
}
