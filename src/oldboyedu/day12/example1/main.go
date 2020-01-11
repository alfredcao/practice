package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_0_0

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("新建生产者失败！err : ", err)
		return
	}
	defer producer.Close()

	for i := 0; i < 100; i++ {
		msg := &sarama.ProducerMessage{
			Topic: "nginx_logs",
			Value: sarama.StringEncoder(fmt.Sprintf("this is a good test, my message is good! %d", i)),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("发送消息失败！err : ", err)
			return
		}

		fmt.Printf("发送消息成功！partition : %d, offset : %d\n", partition, offset)
		time.Sleep(time.Second)
	}

}
