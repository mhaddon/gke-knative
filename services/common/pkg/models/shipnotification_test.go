package models

import (
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestNewShipNotification(t *testing.T) {
	notificationJson := "{ \"registration\": { \"id\": \"1x0d\", \"name\": \"Golden Sausage\", \"captain\": \"Duke Lordship\" }, \"status\": { \"position\": { \"lat\":49.2144, \"long\":2.1312 }, \"velocity\":41 } }"

	newEvent := cloudevents.NewEvent()
	newEvent.SetID(uuid.New().String())

	if err := newEvent.SetData(notificationJson); err != nil {
		t.Error("Error initialising test")
	}

	log.Printf("%v", newEvent)

	notification, err := NewShipNotification(newEvent); if err != nil {
		t.Error("Error normalising event")
	}
	// todo - more extensive testing
	if notification.Registration.Name != "Golden Sausage" {
		t.Error("Failed to convert name")
	}
}
