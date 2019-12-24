package main

import (
	"fmt"
	"math/rand"
	"oldboyedu/day7/example1/balance"
	"time"
)

func main() {
	var servers = make([]*balance.Server, 0)
	for i := 0; i < 16; i++ {
		server := balance.NewServer(fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255)), 8080)
		servers = append(servers, server)
	}

	var balancer balance.Balancer
	balancer = new(balance.RoundRobinBalancer)
	//balancer = new(balance.RandomBalancer)

	for {
		server, err := balancer.DoBalance(servers)
		if err != nil {
			fmt.Println("获取服务器失败！", err)
			break
		} else {
			fmt.Println("获取到服务器", *server)
			time.Sleep(time.Second)
		}
	}
}
