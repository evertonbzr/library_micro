package queue

import (
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type SubscriberQueue struct {
	subject string
	handler func(msg jetstream.Msg)
}

type SubscriberPubSub struct {
	subject string
	handler func(msg *nats.Msg)
}

func NewSubscriberToQueue(subject string, handler func(msg jetstream.Msg)) *SubscriberQueue {
	return &SubscriberQueue{
		subject: subject,
		handler: handler,
	}
}

func NewSubscriberToPubSub(subject string, handler func(msg *nats.Msg)) *SubscriberPubSub {
	return &SubscriberPubSub{
		subject: subject,
		handler: handler,
	}
}

func createSubscriberQueue(subscriber *SubscriberQueue) {
	// SubscriberToQueue(subscriber.subject, subscriber.handler)
}

func createSubscriberPubSub(subscriber *SubscriberPubSub) {
	// SubscriberToPubSub(subscriber.subject, subscriber.handler)
}

func ListenSubscriber(subscribers ...interface{}) {
	for _, subscriber := range subscribers {
		switch s := subscriber.(type) {
		case *SubscriberQueue:
			createSubscriberQueue(s)
			slog.Info("SubscriberQueue", "subject", s.subject)
		case *SubscriberPubSub:
			createSubscriberPubSub(s)
			slog.Info("SubscriberPubSub", "subject", s.subject)
		default:
			panic("Invalid subscriber type")
		}
	}
}
