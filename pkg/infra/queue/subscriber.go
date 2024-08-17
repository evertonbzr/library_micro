package queue

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type SubscriberQueue struct {
	stream  string
	subject string
	handler func(msg jetstream.Msg)
}

type SubscriberPubSub struct {
	subject string
	handler func(msg *nats.Msg)
}

func NewSubscriberToQueue(stream string, subject string, handler func(msg jetstream.Msg)) *SubscriberQueue {
	return &SubscriberQueue{
		subject: subject,
		handler: handler,
		stream:  stream,
	}
}

func NewSubscriberToPubSub(subject string, handler func(msg *nats.Msg)) *SubscriberPubSub {
	return &SubscriberPubSub{
		subject: subject,
		handler: handler,
	}
}

func createSubscriberQueue(subscriber *SubscriberQueue) string {
	ctx := *GetContext()
	js := GetJetStream()

	slug := fmt.Sprintf("%s-%s-%s-worker", streamName, subscriber.stream, subscriber.subject)

	cons, err := js.CreateOrUpdateConsumer(ctx, subscriber.stream, jetstream.ConsumerConfig{
		Durable:       slug,
		Name:          slug,
		FilterSubject: fmt.Sprintf("%s.%s", subscriber.stream, subscriber.subject),
		AckPolicy:     jetstream.AckExplicitPolicy,
	})

	if err != nil {
		slog.Error("Failed to create consumer", "error", err)
	}
	go func() {
		consContext, err := cons.Consume(subscriber.handler)
		if err != nil {
			slog.Error("Failed to consume", "error", err)
			return
		}

		<-ctx.Done()

		slog.Info("Consumer done", "slug", slug)
		consContext.Stop()
	}()

	return slug
}

func createSubscriberPubSub(subscriber *SubscriberPubSub) {
	_, err := natsConnection.Subscribe(subscriber.subject, subscriber.handler)
	if err != nil {
		log.Fatalf("Failed to subscribe to subject %s: %v", subscriber.subject, err)
	}
}

func ListenSubscriber(subscribers ...interface{}) {
	for _, subscriber := range subscribers {
		switch s := subscriber.(type) {
		case *SubscriberQueue:
			slug := createSubscriberQueue(s)
			slog.Info("SubscriberQueue", "slug", slug)
		case *SubscriberPubSub:
			createSubscriberPubSub(s)
			slog.Info("SubscriberPubSub", "subject", s.subject)
		default:
			panic("Invalid subscriber type")
		}
	}
}
