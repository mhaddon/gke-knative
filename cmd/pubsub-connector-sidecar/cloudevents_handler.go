package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go"
	"log"
)

func createCloudEventFromMessage(msg *pubsub.Message) (*cloudevents.Event, error) {
	config := getConfig()

	newEvent := cloudevents.NewEvent()
	newEvent.SetID(msg.ID)
	newEvent.SetSource(fmt.Sprintf("//pubsub.googleapis.com/projects/%v/topics/%v", config.ProjectId, config.TopicName))
	newEvent.SetType("com.google.cloud.pubsub.topic.publish")
	newEvent.SetDataContentType(*cloudevents.StringOfApplicationJSON())

	if err := newEvent.SetData(string(msg.Data)); err != nil {
		return nil, err
	}

	return &newEvent, nil
}

func sendCloudEvent(msg *pubsub.Message, cancel context.CancelFunc) error {
	eventsClient := getEventsClient()

	event, err := createCloudEventFromMessage(msg); if err != nil {
		log.Printf("[PubSub-Connector-Sidecar][CloudEvent][sendCloudEvent] Error creating cloud event: %s\n", err.Error())
		cancel()
	}

	if _, resp, err := eventsClient.Send(context.Background(), *event); err != nil {
		log.Printf("[PubSub-Connector-Sidecar][CloudEvent][sendCloudEvent] Error packaging response: %s\n", err.Error())
		cancel()
		return err
	} else if resp != nil {
		fmt.Printf("Response:\n%s\n", resp)
		fmt.Printf("Got Event Response Context: %+v\n", resp.Context)

		json, err := resp.MarshalJSON(); if err != nil {
			log.Printf("[PubSub-Connector-Sidecar][CloudEvent][sendCloudEvent] Failed to marshal response: %s\n", err.Error())
		}

		if getConfig().EgressTopicName != "none" {
			client := getPubsubClient()

			result := client.egressTopic.Publish(client.context, &pubsub.Message{
				Data: json,
			})

			id, err := result.Get(context.Background()); if err != nil {
				log.Printf("Get: %v", err)

				return err
			}

			log.Printf("Published message with custom attributes; msg ID: %v\n", id)
		}
	}

	return nil
}
