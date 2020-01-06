package ship_notification_connector

import (
	"context"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec/json"
	"github.com/mhaddon/gke-knative/services/common/src/models"
	"github.com/mhaddon/gke-knative/services/common/src/helper"
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

	unpackedEvent, err := helper.UnpackNestedCloudEvent(&event); if err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Handler] Error unpacking event: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	shipNotification := &models.ShipNotification{}

	if err := unpackedEvent.DataAs(shipNotification); err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Handler] Error deserialising event: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	if err := publishNotification(shipNotification); err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Handler] Error publishing notification: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	return nil
}

func getShipServiceUrl() string {
	return fmt.Sprintf("%v%v", getConfig().Services.Ship, "/notifications")
}

func createPutRequest(notification *models.ShipNotification) (*http.Request, error) {
	stringifiedNotification, err := notification.ConvertToJson(); if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPut, getShipServiceUrl(), strings.NewReader(stringifiedNotification)); if err != nil {
		return nil, err
	}

	request.ContentLength = int64(len(stringifiedNotification))

	return request, nil
}

func sendRequest(request *http.Request) error {
	client := &http.Client{}

	response, err := client.Do(request); if err != nil {
		return err
	} else {
		defer response.Body.Close()

		if response.StatusCode < 200 || response.StatusCode >= 400 {
			return fmt.Errorf("bad http response code: %v", response.StatusCode)
		}
	}

	return nil
}

func publishNotification(notification *models.ShipNotification) error {
	request, err := createPutRequest(notification); if err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Publish] Error creating HTTP request: %v\n", err)
	}

	if err := sendRequest(request); err != nil {
		log.Printf("[ShipNotificationConnector][CloudEvent][Publish] Error sending HTTP request: %v\n", err)
	}

	return nil
}