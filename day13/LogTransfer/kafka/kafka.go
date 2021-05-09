package kafka

import (
	"LogTransfer/es"
	"fmt"
	"github.com/Shopify/sarama"
)

// 初始化kafka消费者 从kafka取数据发往ES

// Init ...
func Init(addr []string, topic string) error {
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Println("failed to start consumer, err:", err)
		return err
	}
	partitions, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Println("failed to get list of partition:", err)
		return err
	}
	fmt.Println(partitions)
	for partition := range partitions { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err:%v\n", partition, err)
			return err
		}
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 直接发给ES
				ld := &es.LogData{
					Topic: topic,
					Data:  string(msg.Value),
				}
				// err = es.SendToES(topic, ip, ld) // 函数调函数
				// 优化一下：直接放到一个channel中
				es.SendToESChan(ld)
			}
		}(pc)
		// pc.AsyncClose()
	}
	return err
}
