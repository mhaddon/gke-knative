package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func pull(onReceive func(msg *pubsub.Message, cancel context.CancelFunc) error) error {
	client := getPubsubClient()

	return client.subscription.Receive(client.contextWithCancel, func(ctx context.Context, msg *pubsub.Message) {
		client.lock.Lock()
		msg.Ack()
		defer client.lock.Unlock()
		if err := onReceive(msg, client.cancel); err != nil {
			log.Printf("[PubSub-Connector-Sidecar][CloudEvent][sendCloudEvent] Error processing messages: %v\n", err.Error())
		}
	})
}
