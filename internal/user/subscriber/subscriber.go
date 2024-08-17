package subscriber

import (
	"fmt"

	"github.com/evertonbzr/library_micro/pkg/infra/queue"
	"github.com/nats-io/nats.go"
)

func CreateBookSubscriber() *queue.SubscriberPubSub {
	return queue.NewSubscriberToPubSub("module_book.create", func(msg *nats.Msg) {
		fmt.Println("BookSubscriber")
		msg.Ack()
	})
}

func UpdateBookSubscriber() *queue.SubscriberPubSub {
	return queue.NewSubscriberToPubSub("module_book.update", func(msg *nats.Msg) {
		fmt.Println("BookSubscriber")
		msg.Ack()
	})
}

func GetAll() []interface{} {
	return []interface{}{
		CreateBookSubscriber(),
		UpdateBookSubscriber(),
	}
}
