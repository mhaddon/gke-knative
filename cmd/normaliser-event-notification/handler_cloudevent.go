package main

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/google/uuid"
	"github.com/mhaddon/gke-knative/pkg/models"
	"log"
)

func publishNormalisedMessage(shipNotification models.ShipNotification, response *cloudevents.EventResponse) error {
	newEvent := cloudevents.NewEvent()
	newEvent.SetID(uuid.New().String())
	newEvent.SetSource(getConfig().EgressSource)
	newEvent.SetType(getConfig().EgressType)

	jsonShipNotification, err := shipNotification.ConvertToJson(); if err != nil {
		log.Printf("[Normaliser-Event-Normaliser][CloudEvent][publishNormalisedMessage] Error serialising response: %s\n", err.Error())
		return err
	}

	if err := newEvent.SetData(jsonShipNotification); err != nil {
		log.Printf("[Normaliser-Event-Normaliser][CloudEvent][publishNormalisedMessage] Error packaging response: %s\n", err.Error())
		return err
	}

	newEvent.SetDataContentType(*cloudevents.StringOfApplicationJSON())
	response.RespondWith(200, &newEvent)

	return nil
}

func cloudWatchHandler(ctx context.Context, event cloudevents.Event, rawEvent cloudevents.Event, response *cloudevents.EventResponse) error {
	normalisedShipNotification, err := normaliseEvent(event); if err != nil {
		log.Printf("[Normaliser-Event-Normaliser][CloudEvent][Handler] Error normalising event: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	if err := publishNormalisedMessage(*normalisedShipNotification, response); err != nil {
		log.Printf("[Normaliser-Event-Normaliser][CloudEvent][Handler] Error Publishing Response: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	return nil
}
