package tools

import (
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

/*
var (
	wg sync.WaitGroup
)
*/

type KafkaClientSt struct {
	Client sarama.Consumer
	Addr   string
	Topic  string
	Wg     sync.WaitGroup
}

var (
	KafkaClient *KafkaClientSt
)

func InitKafka(addr string, topic string) (err error) {

	KafkaClient = &KafkaClientSt{}

	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	KafkaClient.Client = consumer
	KafkaClient.Addr = addr
	KafkaClient.Topic = topic
	return

	/*
		partitionList, err := consumer.Partitions(topic)
		if err != nil {
			logs.Error("Failed to get the list of partitions: ", err)
			return
		}

		for partition := range partitionList {
			pc, errRet := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
			if errRet != nil {
				err = errRet
				logs.Error("Failed to start consumer for partition %d: %s\n", partition, err)
				return
			}
			defer pc.AsyncClose()
			go func(pc sarama.PartitionConsumer) {
				//wg.Add(1)
				for msg := range pc.Messages() {
					logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
					//fmt.Println()
				}
				//wg.Done()
			}(pc)
		}*/
	//time.Sleep(time.Hour)
	//wg.Wait()
	//consumer.Close()
	//return
}
