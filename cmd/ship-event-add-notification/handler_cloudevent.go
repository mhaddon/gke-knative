package main

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/mhaddon/gke-knative/services/common/pkg/models"
	"log"
)

func cloudWatchHandler(ctx context.Context, event cloudevents.Event, rawEvent cloudevents.Event, response *cloudevents.EventResponse) error {
	shipNotification := &models.ShipNotification{}

	if err := event.DataAs(shipNotification); err != nil {
		log.Printf("[Ship-Event-Add-Notification][CloudEvent][Handler] Error deserialising event: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	if err := addNotification(shipNotification); err != nil {
		log.Printf("[Ship-Event-Add-Notification][CloudEvent][Handler] Error publishing notification: %s, ID: %v\n", err.Error(), event.Context.GetID())
		return err
	}

	return nil
}

func addNotification(notification *models.ShipNotification) error {
	if err := getPersistence().GetCollection().Insert(&notification); err != nil {
		log.Printf("[Ship-Event-Add-Notification][DAO] Error saving ship notification... %v", err)
		return err
	}
	return nil
}