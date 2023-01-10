package rabbitmq

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Exchange struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type BindingOptions struct {
	RoutingKey string
	NoWait     bool
	Args       amqp.Table
}

type ConsumerOptions struct {
	Tag       string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

type Session struct {
	Exchange        Exchange
	Queue           Queue
	BindingOptions  BindingOptions
	ConsumerOptions ConsumerOptions
}

type RabbitMQ struct {
	conn   *amqp.Connection
	notify chan error
}

func NewConnection(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	rabbitMQ := &RabbitMQ{
		conn:   conn,
		notify: make(chan error, 1),
	}

	go func() {
		for err := range conn.NotifyClose(make(chan *amqp.Error)) {
			rabbitMQ.notify <- errors.New(err.Error())
		}
	}()

	return rabbitMQ, nil
}

func (rmq *RabbitMQ) Connection() *amqp.Connection {
	return rmq.conn
}

func (rmq *RabbitMQ) Notify() <-chan error {
	return rmq.notify
}

func (rmq *RabbitMQ) Shutdown() error {
	if rmq.conn.IsClosed() {
		return nil
	}

	return rmq.conn.Close()
}
