package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: time.Second * 3,
	})

	if err != nil {
		fmt.Println("connect etcd failed, err :", err)
		return
	}

	defer cli.Close()
	fmt.Println("connect success")
}
