package main

import (
	"fmt"
	kafka2 "github.com/loressl/imersao_full_stack_cycle/a´pp/kafka"
	"github.com/loressl/imersao_full_stack_cycle/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume() // thread fica ouvindo, se não receber nada fica parada

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg) // cada produção pode rodar de forma concorrente com o go, trabalhando em paralelo
	}
}