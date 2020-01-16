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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	watchChan := cli.Watch(ctx, "watchKey")
	for watchResponse := range watchChan {
		for _, event := range watchResponse.Events {
			fmt.Printf("receive event, type : %v, %s=%s\n", event.Type, string(event.Kv.Key), string(event.Kv.Value))
		}
	}
}
