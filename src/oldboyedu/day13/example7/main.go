package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

const topic = "nginx_logs"

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, sarama.NewConfig())
	if err != nil {
		fmt.Println("new consumer failed, err :", err)
		return
	}
	defer consumer.Close()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("get partitions failed, err :", err)
		return
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)

		if err != nil {
			fmt.Printf("consume partition[%d] failed, err %v:\v", partition, err)
			continue
		}

		defer pc.AsyncClose()

		go func() {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d, offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}()

	}

	time.Sleep(time.Hour)

}
