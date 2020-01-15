package main

import (
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	_, err = cli.Put(ctx, "/logagent/conf", "confValue")
	cancel()
	if err != nil {
		fmt.Println("put failed, err :", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	getRes, err := cli.Get(ctx, "/logagent/conf")
	cancel()
	if err != nil {
		fmt.Println("get failed, err :", err)
		return
	}

	for _, v := range getRes.Kvs {
		fmt.Printf("get key %s value is %s\n", string(v.Key), string(v.Value))
	}

}
