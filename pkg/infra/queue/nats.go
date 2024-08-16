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

func CloseNatsConnection() jetstream.JetStream {
	if natsConnection != nil {
		natsConnection.Close()
		natsConnection = nil
		natsJetStream = nil
	}

	return natsJetStream
}
