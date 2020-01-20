package main

import (
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	"sync"
)

var (
	eventsClientInstance client.Client
	eventsClientOnce     sync.Once
)

func getEventsClient() client.Client {
	eventsClientOnce.Do(func() {
		config := getConfig()

		t, err := cloudevents.NewHTTPTransport(
			cloudevents.WithTarget(config.TargetUrl),
			cloudevents.WithStructuredEncoding(),
		)

		c, err := cloudevents.NewClient(t); if err != nil {
			panic("unable to create cloudevent client: " + err.Error())
		}

		eventsClientInstance = c
	})

	return eventsClientInstance
}



