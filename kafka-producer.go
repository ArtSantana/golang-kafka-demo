package main

import (
	"math/rand"
	"strconv"

	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)


func produces(ctx context.Context) {
    kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{brokerAddress},
        Topic: topic,
    })

    err := kafkaWriter.WriteMessages(ctx, kafka.Message{
        Key: []byte(strconv.Itoa(rand.Intn(100))),
        Value: []byte("This message in being sended by a lime gin tonic"),
    })

    if err != nil {
        panic("could not write message " + err.Error())
    }
}
