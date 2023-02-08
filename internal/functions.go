package internal

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
	"nats-publisher/configs"
	"nats-publisher/domain"
)

type NatsQueue struct {
	conn stan.Conn
	subj string
}

type OrdersMessageQueue interface {
	SendNewMessage(msg []byte) error
	Close()
}

func NewNatsQueue(config *configs.Config) (*NatsQueue, error) {
	opts := []nats.Option{
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			log.Fatalf("conn error: %s", err.Error())
		}),
	}

	conn, err := nats.Connect(config.URL, opts...)
	if err != nil {
		return nil, err
	}

	sc, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsConn(conn))
	if err != nil {
		return nil, err
	}

	return &NatsQueue{
		conn: sc,
		subj: config.Subj,
	}, nil
}

func (nq *NatsQueue) SendNewMessage(msg []byte) error {
	err := nq.conn.Publish(nq.subj, msg)
	if err != nil {
	}
	return nil
}

func (nq *NatsQueue) Close() {
	nq.conn.Close()
}

func CreateNewOrder() ([]byte, error) {
	var order = domain.Order{}
	gofakeit.Seed(0)
	gofakeit.Struct(&order)

	jOrder, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	return jOrder, nil
}
