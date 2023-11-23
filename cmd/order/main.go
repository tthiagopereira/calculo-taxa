package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"gitub.com/thiagopereira7/taxas/internal/infra/database"
	"gitub.com/thiagopereira7/taxas/internal/usecase"
	"gitub.com/thiagopereira7/taxas/pkg/rabbitmq"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChanel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, msgRabbitmqChannel)

	rabbitmqWorker(msgRabbitmqChannel, uc)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Menssagem processada e salva no banco:", output)
	}
}
