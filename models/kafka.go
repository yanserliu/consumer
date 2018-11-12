package models

import (
	"consumer/infrastructure/db"
	"fmt"
	"math/rand"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
)

var (
	wg sync.WaitGroup
)

type update struct {
	Data      []data  `json:"data"`
	Database  string  `json:"database"`
	Es        int64   `json:"es"`
	Id        int     `json:"id"`
	IsDdl     bool    `json:"isDdl"`
	MysqlType data    `json:"mysqlType"`
	Old       []data  `json:"old"`
	Sql       string  `json:"sql"`
	SqlType   sqlType `json:"sqlType"`
	Table     string  `json:"table"`
	Ts        int64   `json:"ts"`
	Type      string  `json:"type"`
}

type data struct {
	Id       string `json:"id"`
	User     string `json:"user"`
	Sn       string `json:"sn"`
	Location string `json:"location"`
}

type sqlType struct {
	Id       int `json:"id"`
	User     int `json:"user"`
	Sn       int `json:"sn"`
	Location int `json:"location"`
}

func Consumer() {

	kafka_addr := beego.AppConfig.String("kafka::addr")
	topic := beego.AppConfig.String("kafka::topic")
	var user db.User
	consumer, err := sarama.NewConsumer([]string{kafka_addr}, nil)
	if err != nil {
		panic(err)
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("example", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {

				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				user.Email = string(msg.Value)
				user.Id = rand.Int63n(1000)
				_, err = db.Insert(&user)
				if err != nil {
					fmt.Printf("insert err:", err)
				}

			}
		}(pc)
		wg.Wait()
		consumer.Close()
	}

}
