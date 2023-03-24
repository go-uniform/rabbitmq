package rabbitmq

import "log"

func (r *rabbitmq) Pop() {

	q, err := r.Channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		panic(err)
	}

	messages, err := r.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	for message := range messages {
		log.Printf("Received a message: %s", message.Body)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
