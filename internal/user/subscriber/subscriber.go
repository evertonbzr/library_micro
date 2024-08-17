package subscriber

// func V1DeleteBookSubscriber() *queue.SubscriberQueue {
// 	return queue.NewSubscriberToQueue("module_books", "v1-delete", func(msg jetstream.Msg) {
// 		fmt.Println("BookSubscriber")
// 		// msg.Ack()
// 		msg.Ack()
// 	})
// }

func GetAll() []interface{} {
	return []interface{}{}
}
