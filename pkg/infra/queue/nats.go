package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var natsConnection *nats.Conn
var natsJetStream jetstream.JetStream
var streamName string

func ConnectNats(ctx context.Context, uri string, name string) error {
	if natsConnection != nil {
		return nil
	}

	conn, err := nats.Connect(uri)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
		return err
	}

	js, err := jetstream.New(conn)
	if err != nil {
		log.Fatalf("Failed to connect to JetStream: %v", err)
		return err
	}

	natsConnection = conn
	natsJetStream = js
	streamName = name

	natsJetStream.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     name,
		Subjects: []string{fmt.Sprintf("%s.*", name)},
	})

	return nil
}

func GetNatsConnection() *nats.Conn {
	if natsConnection == nil {
		log.Fatalf("NATS connection is not initialized")
	}

	return natsConnection
}

func GetJetStream() jetstream.JetStream {
	if natsJetStream == nil {
		log.Fatalf("JetStream is not initialized")
	}

	return natsJetStream
}

func CloseNatsConnection() jetstream.JetStream {
	if natsConnection != nil {
		natsConnection.Close()
		natsConnection = nil
		natsJetStream = nil
	}

	return natsJetStream
}

func SubscriberToPubSub(subject string, callback func(msg *nats.Msg)) {
	_, err := GetNatsConnection().Subscribe(subject, callback)
	if err != nil {
		log.Fatalf("Failed to subscribe to subject %s: %v", subject, err)
	}
}

func SubscribeToQueue(ctx context.Context, stream string, subject string, handler jetstream.MessageHandler) {
	slug := fmt.Sprintf("%s-%s-%s-worker", streamName, stream, subject)

	cons, _ := natsJetStream.CreateOrUpdateConsumer(ctx, stream, jetstream.ConsumerConfig{
		AckPolicy:     jetstream.AckExplicitPolicy,
		Durable:       slug,
		Name:          slug,
		FilterSubject: fmt.Sprintf("%s.%s", stream, subject),
	})

	consContext, _ := cons.Consume(handler)

	defer consContext.Closed()
	defer slog.Info("Consumer closed")
}

func DecodeMessage(msg *nats.Msg, v interface{}) error {
	if err := json.Unmarshal(msg.Data, v); err != nil {
		log.Printf("Failed to decode message: %v", err)
		return err
	}

	return nil
}

func EncodeMessage(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("Failed to encode message: %v", err)
		return nil
	}

	return data
}
