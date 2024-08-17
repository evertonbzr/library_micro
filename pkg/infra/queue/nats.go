package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var natsConnection *nats.Conn
var natsJetStream jetstream.JetStream
var streamName string
var contextStream *context.Context

func ConnectNats(ctx *context.Context, uri string, name string) error {
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
	contextStream = ctx

	natsJetStream.CreateOrUpdateStream(*ctx, jetstream.StreamConfig{
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

func GetContext() *context.Context {
	if contextStream == nil {
		log.Fatalf("Context is not initialized")
	}

	return contextStream
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
