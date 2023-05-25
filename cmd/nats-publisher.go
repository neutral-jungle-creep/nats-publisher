package main

import (
	"log"
	"nats-publisher/configs"
	"nats-publisher/internal"
	"time"
)

func main() {
	config := configs.NewConfig()

	natsQueue, err := internal.NewNatsQueue(config)
	if err != nil {
		log.Fatalf("can't connect to nats: %s", err.Error())
	}
	log.Printf("connect to nats is successfully: url = [%s], cluster = [%s], client = [%s]",
		config.URL, config.ClusterID, config.ClientID)

	var messageQueue internal.OrdersMessageQueue = natsQueue
	defer messageQueue.Close()

	for {
		order, _ := internal.CreateNewOrder()

		if err := messageQueue.SendNewMessage(order); err != nil {
			log.Printf("error: %s", err.Error())
		}
		log.Printf("Published [%s] : '%s'\n", config.Subj, order)
		time.Sleep(config.Sleep * time.Second)
	}

}
