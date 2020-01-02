package ship_notification_connector

import (
	"context"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec/json"
	"github.com/mhaddon/gke-knative/services/common/src/models"
	"log"
	"net/http"
	"strings"
)

func CreateCloudWatchListener() error {
	config := getConfig()

	t, err := cloudevents.NewHTTPTransport(cloudevents.WithPort(config.CloudEvent.Port)); if err != nil {
		log.Fatalf("[ShipNotificationConnector][CloudEvent] Failed to create HTTP transport layer, %v", err)
		return err
	}

	c, err := cloudevents.NewClient(t); if err != nil {
		log.Fatalf("[ShipNotificationConnector][CloudEvent] Failed to create CloudEvent client, %v", err)
		return err
	}

	configureCodec()

	log.Printf("[ShipNotificationConnector][CloudEvent] Listening for cloud events on port: %v...", config.CloudEvent.Port)

	return c.StartReceiver(context.Background(), cloudWatchHandler)
}

//todo - replace with solution from resolution of this ticket: https://github.com/cloudevents/sdk-go/issues/254
func configureCodec() {
	log.Print("[ShipNotificationConnector][CloudEvent][ConfigureCodec] Configuring 'application/octet-stream' content type.")

	datacodec.AddDecoder("application/octet-stream", json.Decode)
	datacodec.AddEncoder("application/octet-stream", json.Encode)
}

func cloudWatchHandler(ctx context.Context, event cloudevents.Event, response *cloudevents.EventResponse) error {
	log.Printf("[ShipNotificationConnector][CloudEvent][Handler] Recieved Event with ID: %v, Source: %v, Subject: %v, Type: %v, Time: %v", event.Context.GetID(), event.Context.GetSource(), event.Context.GetSubject(), event.Context.GetType(), event.Context.GetTime())

	shipNotification := &models.ShipNotification{}

	if err := event.DataAs(shipNotification); err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Handler] Error deserialising event: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	if err := publishNotification(shipNotification); err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Handler] Error publishing notification: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	return nil
}

func publishNotification(notification *models.ShipNotification) error {
	client := &http.Client{}

	stringifiedNotification, err := notification.ConvertToJson(); if err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Publish] Error serialising notification: %v\n", err)
		return err
	}

	request, err := http.NewRequest("PUT", fmt.Sprintf("%v%v", getConfig().Services.Ship, "/notifications"), strings.NewReader(stringifiedNotification)); if err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Publish] Error creating HTTP request: %v\n", err)
		return err
	}

	request.ContentLength = len(stringifiedNotification)

	response, err := client.Do(request); if err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Publish] Error sending HTTP request: %v\n", err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 400 {
		log.Printf("[ShipNotificationConnector][CloudEvent][Publish] Bad HTTP response code: %v\n", response.StatusCode)
		return err
	}

	return nil
}
