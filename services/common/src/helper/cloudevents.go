package helper

import (
	"log"
	cloudevents "github.com/cloudevents/sdk-go"
)

type NestedCloudEventT struct {
	CloudEventsVersion string `json:"cloudEventsVersion"`
}

type NestedCloudEvent struct {
	CloudEventsVersion string `json:"cloudEventsVersion"`
	ContentType string `json:"contentType"`
	Data interface{} `json:"data"`
	EventID string `json:"eventID"`
	EventTime string `json:"eventTime"`
	EventType string `json:"eventType"`
	Extensions interface{} `json:"extensions"`
	Source string `json:"source"`
}

func UnpackNestedCloudEvent(event *cloudevents.Event) (*cloudevents.Event, error) {
	nestedCloudEvent := &NestedCloudEvent{}

	if err := event.DataAs(nestedCloudEvent); err != nil {
		return event, nil
	}

	if len(nestedCloudEvent.CloudEventsVersion) ==  0 {
		return event, nil
	}

	newEvent := cloudevents.NewEvent()
	newEvent.SetID(nestedCloudEvent.EventID)
	newEvent.SetSource(nestedCloudEvent.Source)
	newEvent.SetType(nestedCloudEvent.EventType)
	if err := newEvent.SetData(nestedCloudEvent.Data); err != nil {
		log.Printf("[CloudEvent][UnpackNestedCloudEvent] Error unpacking event: %s\n", err.Error())
		return nil, err
	}
	newEvent.SetDataContentType(nestedCloudEvent.ContentType)

	return &newEvent, nil
}