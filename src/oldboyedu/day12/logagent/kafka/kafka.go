package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"oldboyedu/day12/logagent/conf"
	"oldboyedu/day12/logagent/es"
	"oldboyedu/day12/logagent/tailf"
	"time"
)

var CONSUME_TOPIC_LIST = []string{"nginx_log"}

func InitKafkaProducer() (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_0_0

	producer, producerErr := sarama.NewSyncProducer([]string{conf.AppConfig.KafkaConf.Addr}, config)
	if producerErr != nil {
		err = fmt.Errorf("new kafka producer failed, err : %v", err)
		return
	}

	go func() {
		for {
			tailMsg := <-tailf.TailObjManager.MsgChan

			msg := &sarama.ProducerMessage{
				Topic: tailMsg.Topic,
				Value: sarama.StringEncoder(tailMsg.Text),
			}

			partition, offset, err := producer.SendMessage(msg)
			if err != nil {
				logs.Error("send msg to kafka failed, err : ", err)
				return
			}

			logs.Info("send msg to kafka success！topic : %s, partition : %d, offset : %d, msg : %s", msg.Topic, partition, offset, msg.Value)

		}
	}()

	return
}

func InitKafkaConsumer() (err error) {
	consumer, err := sarama.NewConsumer([]string{conf.AppConfig.KafkaConf.Addr}, sarama.NewConfig())
	if err != nil {
		logs.Error("new kafka consumer failed, err : ", err)
		return
	}

	for _, topic := range CONSUME_TOPIC_LIST {
		go func() {
			partitions, err := consumer.Partitions(topic)
			if err != nil {
				logs.Error("get topic[%s] partitions failed, err : %v", topic, err)
				return
			} else {
				logs.Info("get topic[%s] partitions success, partitions : %v", topic, partitions)
			}

			for _, partition := range partitions {
				pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)

				if err != nil {
					logs.Error("consume topic[%s] partition[%d] failed, err : %v:", topic, partition, err)
					continue
				} else {
					logs.Info("consume topic[%s] partition[%d] success", topic, partition)
				}

				go func() {
					defer pc.AsyncClose()
					for msg := range pc.Messages() {
						value := string(msg.Value)
						//logs.Info("consume message topic[%s] partition[%d] success, offset -> %d, key -> %s, value -> %s",
						//	topic, partition, msg.Partition, msg.Offset, key, value)
						esDoc := ESDoc{
							Msg: value,
						}

						ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
						_, err = es.ESClient.Index().
							Index(topic).
							Type(topic).
							BodyJson(esDoc).
							Do(ctx)
						cancel()
						if err != nil {
							logs.Error("put document to es failed, err :%v", err)
							continue
						}

						logs.Info("put document to es success")

					}
				}()

			}
		}()
	}

	return
}

type ESDoc struct {
	Msg string
}
