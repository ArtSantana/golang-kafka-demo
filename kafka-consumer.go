package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func consumes(ctx context.Context) {
    kafkaReader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{brokerAddress},
        Topic: topic,
        
    })

	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}}
