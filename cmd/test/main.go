package main

import (
	kafka2 "miniTikTok/cmd/video/kafka"
	"time"
)

func main() {
	go func() {
		Put()
	}()
	go func() {
		for {
			Get()
		}
	}()

	for {
		time.Sleep(time.Hour * 60)
	}
}

func Put() {
	_, err := kafka2.Send("example", "777")
	if err != nil {
		return
	}
}

func Get() {
	c := new(kafka2.KafkaConsumer)
	c.Node = []string{"127.0.0.1:9092"}
	c.Topic = "example"
	c.Consume()
	//time.Sleep(1 * time.Second)
}
