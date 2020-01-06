package normaliser

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec/json"
	"github.com/google/uuid"
	"github.com/mhaddon/gke-knative/services/common/src/models"
	"github.com/mhaddon/gke-knative/services/common/src/helper"
	"log"
)

func CreateCloudWatchListener() error {
	config := getConfig()

	t, err := cloudevents.NewHTTPTransport(cloudevents.WithPort(config.CloudEvent.Port)); if err != nil {
		log.Fatalf("[Normaliser][CloudEvent] Failed to create HTTP transport layer, %v", err)
		return err
	}

	c, err := cloudevents.NewClient(t); if err != nil {
		log.Fatalf("[Normaliser][CloudEvent] Failed to create CloudEvent client, %v", err)
		return err
	}

	configureCodec()

	log.Printf("[Normaliser][CloudEvent] Listening for cloud events on port: %v...", config.CloudEvent.Port)

	return c.StartReceiver(context.Background(), cloudWatchHandler)
}

//todo - replace with solution from resolution of this ticket: https://github.com/cloudevents/sdk-go/issues/254
func configureCodec() {
	log.Print("[Normaliser][CloudEvent][ConfigureCodec] Configuring 'application/octet-stream' content type.")

	datacodec.AddDecoder("application/octet-stream", json.Decode)
	datacodec.AddEncoder("application/octet-stream", json.Encode)
}

func publishNormalisedMessage(shipNotification models.ShipNotification, response *cloudevents.EventResponse) error {
	newEvent := cloudevents.NewEvent()
	newEvent.SetID(uuid.New().String())
	newEvent.SetSource(getConfig().Egress.Source)
	newEvent.SetType(getConfig().Egress.Type)

	jsonShipNotification, err := shipNotification.ConvertToJson(); if err != nil {
		log.Printf("[Normaliser][CloudEvent][publishNormalisedMessage] Error serialising response: %s\n", err.Error())
		return err
	}

	if err := newEvent.SetData(jsonShipNotification); err != nil {
		log.Printf("[Normaliser][CloudEvent][publishNormalisedMessage] Error packaging response: %s\n", err.Error())
		return err
	}

	newEvent.SetDataContentType(*cloudevents.StringOfApplicationJSON())
	response.RespondWith(200, &newEvent)

	return nil
}

func cloudWatchHandler(ctx context.Context, event cloudevents.Event, response *cloudevents.EventResponse) error {
	log.Printf("[Normaliser][CloudEvent][Handler] Recieved Event with ID: %v, Source: %v, Subject: %v, Type: %v, Time: %v", event.Context.GetID(), event.Context.GetSource(), event.Context.GetSubject(), event.Context.GetType(), event.Context.GetTime())

	unpackedEvent, err := helper.UnpackNestedCloudEvent(&event); if err != nil {
		log.Printf("[Normaliser][CloudEvent][Handler] Error unpacking event: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	normalisedShipNotification, err := normaliseEvent(*unpackedEvent); if err != nil {
		log.Printf("[Normaliser][CloudEvent][Handler] Error normalising event: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	if err := publishNormalisedMessage(*normalisedShipNotification, response); err != nil {
		log.Printf("[Normaliser][CloudEvent][Handler] Error Publishing Response: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	return nil
}
