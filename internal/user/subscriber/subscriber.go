package subscriber

import (
	"fmt"

	"github.com/evertonbzr/library_micro/pkg/infra/queue"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func V1CreateBookSubscriber() *queue.SubscriberPubSub {
	return queue.NewSubscriberToPubSub("module_book.v1-create", func(msg *nats.Msg) {
		fmt.Println("BookSubscriber")
		msg.Ack()
	})
}

func V1UpdateBookSubscriber() *queue.SubscriberPubSub {
	return queue.NewSubscriberToPubSub("module_book.v1-update", func(msg *nats.Msg) {
		fmt.Println("BookSubscriber")
		msg.Ack()
	})
}

func V1DeleteBookSubscriber() *queue.SubscriberQueue {
	return queue.NewSubscriberToQueue("module_books", "v1-delete", func(msg jetstream.Msg) {
		fmt.Println("BookSubscriber")
		// msg.Ack()
		msg.Ack()
	})
}

func GetAll() []interface{} {
	return []interface{}{
		V1CreateBookSubscriber(),
		V1UpdateBookSubscriber(),
		V1DeleteBookSubscriber(),
	}
}
