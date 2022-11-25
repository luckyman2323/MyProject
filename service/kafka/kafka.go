/*
 * @Author: liyj liyj@yunphant.com
 * @Date: 2022-09-16 09:35:12
 * @LastEditors: liyj liyj@yunphant.com
 * @LastEditTime: 2022-09-16 09:46:17
 * @FilePath: /myproject/service/kafka.go
 * @Description:
 *
 * Copyright (c) 2022 by liyj liyj@yunphant.com, All Rights Reserved.
 */
package kafka

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

func SendMessageToKafka() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，我们默认设置32个分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "task"
	msg.Value = sarama.StringEncoder("producer kafka messages...")

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.20.216.8:9092"}, config)
	if err != nil {
		fmt.Println("Producer closed, err:", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

func ConsumerMessageFromKafka() {
	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{"192.20.216.8:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	partitionList, err := consumer.Partitions("task-status-data") // 通过topic获取到所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList { // 遍历所有的分区
		pc, err := consumer.ConsumePartition("task", int32(partition), sarama.OffsetNewest) // 针对每个分区创建一个分区消费者
		if err != nil {
			fmt.Println("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { // 为每个分区开一个go协程取值
			for msg := range pc.Messages() { // 阻塞直到有值发送过来，然后再继续等待
				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
