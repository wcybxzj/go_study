package server

import (
	"13.oldboy/project/log_transfer/tools"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

func Run() (err error) {

	partitionList, err := tools.KafkaClient.Client.Partitions(tools.KafkaClient.Topic)
	if err != nil {
		logs.Error("Failed to get the list of partitions: ", err)
		return
	}

	for partition := range partitionList {
		pc, errRet := tools.KafkaClient.Client.ConsumePartition(
			tools.KafkaClient.Topic, int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			logs.Error("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			tools.KafkaClient.Wg.Add(1)
			for msg := range pc.Messages() {
				logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				//fmt.Println()
				err = tools.SendToES(tools.KafkaClient.Topic, msg.Value)
				if err != nil {
					logs.Warn("send to es failed, err:%v", err)
				}
			}
			tools.KafkaClient.Wg.Done()
		}(pc)
	}

	tools.KafkaClient.Wg.Wait()
	return
}
