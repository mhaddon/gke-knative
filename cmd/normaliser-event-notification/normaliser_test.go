package main

import (
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go"
	"testing"
)

func TestGetVersion(t *testing.T) {
	expectedResult := 2

	newEvent := cloudevents.NewEvent()

	if err := newEvent.SetData(fmt.Sprintf("{ \"version\": %v }", expectedResult)); err != nil {
		t.Error("Error initialising test")
	}

	version, err := getVersion(newEvent); if err != nil {
		t.Error("Error getting version")
	}

	if version != expectedResult {
		t.Errorf("Invalid version, expected: %v, got: %v", expectedResult, version)
	}
}

func TestGetVersionEmpty(t *testing.T) {
	expectedResult := 0

	newEvent := cloudevents.NewEvent()

	if err := newEvent.SetData(fmt.Sprintf("{ \"cat\": %v }", expectedResult)); err != nil {
		t.Error("Error initialising test")
	}

	_, err := getVersion(newEvent); if err == nil {
		t.Error("Expected error response")
	}
}

func TestVersionOneConvert(t *testing.T) {
	versionOneJson := "{ \"version\": 1, \"id\": \"1x0d\", \"name\": \"Golden Sausage\", \"captain\": \"Duke Lordship\", \"position\": \"49.2144,2.1312\", \"speed\": { \"velocity\": 40, \"unit\": \"Knots\" } }"

	newEvent := cloudevents.NewEvent()

	if err := newEvent.SetData(versionOneJson); err != nil {
		t.Error("Error initialising test")
	}

	notification, err := normaliseEvent(newEvent); if err != nil {
		t.Error("Error normalising event")
	}

	// todo - more extensive testing
	if notification.Registration.Name != "Golden Sausage" {
		t.Error("Failed to convert name")
	}
}

func TestVersionTwoConvert(t *testing.T) {
	versionTwoJson := "{ \"version\": 2, \"ship\": { \"registration\": \"1x0d\", \"name\": \"Golden Sausage\" }, \"captain\": \"Duke Lordship\", \"position\": { \"latitude\": 49.2144, \"longitude\": 2.1312 }, \"speed\": 41 }"

	newEvent := cloudevents.NewEvent()

	if err := newEvent.SetData(versionTwoJson); err != nil {
		t.Error("Error initialising test")
	}

	notification, err := normaliseEvent(newEvent); if err != nil {
		t.Error("Error normalising event")
	}

	// todo - more extensive testing
	if notification.Registration.Name != "Golden Sausage" {
		t.Error("Failed to convert name")
	}
}
