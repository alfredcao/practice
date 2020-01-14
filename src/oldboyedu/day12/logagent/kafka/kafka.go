package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"oldboyedu/day12/logagent/conf"
	"oldboyedu/day12/logagent/tailf"
)

func InitKafka() (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_0_0

	producer, producerErr := sarama.NewSyncProducer([]string{conf.AppConfig.KafkaConf.Addr}, config)
	if producerErr != nil {
		err = fmt.Errorf("init kafka producer failed, err : %v", err)
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

			logs.Info("send msg to kafka success！partition : %d, offset : %d", partition, offset)
		}
	}()

	return
}
