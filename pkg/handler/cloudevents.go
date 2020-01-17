package handler

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec/json"
	"github.com/mhaddon/gke-knative/pkg/helper"
	"log"
	"time"
)

//todo - replace with solution from resolution of this ticket: https://github.com/cloudevents/sdk-go/issues/254
func configureCodec() {
	log.Print("[Ship-Event-Add-Notification][CloudEvent][ConfigureCodec] Configuring 'application/octet-stream' content type.")

	datacodec.AddDecoder("application/octet-stream", json.Decode)
	datacodec.AddEncoder("application/octet-stream", json.Encode)
}

func unpackEvent(callback func(ctx context.Context, unpackedEvent cloudevents.Event, rawEvent cloudevents.Event, response *cloudevents.EventResponse) error) func(ctx context.Context, event cloudevents.Event, response *cloudevents.EventResponse) error {
	return func(ctx context.Context, event cloudevents.Event, response *cloudevents.EventResponse) error {
		log.Printf("[CloudEvent][Handler] Recieved Event with ID: %v, Source: %v, Subject: %v, Type: %v, Time: %v", event.Context.GetID(), event.Context.GetSource(), event.Context.GetSubject(), event.Context.GetType(), event.Context.GetTime())
		startNanoSeconds := time.Now().UnixNano()

		unpackedEvent, err := helper.UnpackNestedCloudEvent(&event); if err != nil {
			log.Printf("[CloudEvent][Handler] ERROR: Event with ID: %v, Source: %v, Subject: %v, Type: %v, Time: %v has resulted in an error: %v", event.Context.GetID(), event.Context.GetSource(), event.Context.GetSubject(), event.Context.GetType(), event.Context.GetTime(), err)
			return err
		}

		if err := callback(ctx, *unpackedEvent, event, response); err != nil {
			log.Printf("[CloudEvent][Handler] ERROR: Event with ID: %v, Source: %v, Subject: %v, Type: %v, Time: %v has resulted in an error: %v", event.Context.GetID(), event.Context.GetSource(), event.Context.GetSubject(), event.Context.GetType(), event.Context.GetTime(), err)
		}

		durationNanoseconds := time.Now().UnixNano() - startNanoSeconds

		log.Printf("[CloudEvent][Handler] Event with ID: %v, Source: %v, Subject: %v, Type: %v, Time: %v, took %v nanoseconds or %v milliseconds", event.Context.GetID(), event.Context.GetSource(), event.Context.GetSubject(), event.Context.GetType(), event.Context.GetTime(), durationNanoseconds, durationNanoseconds / int64(time.Millisecond))

		return nil
	}
}

func CreateCloudWatchListener(port int, handler func(ctx context.Context, unpackedEvent cloudevents.Event, rawEvent cloudevents.Event, response *cloudevents.EventResponse) error) error {
	t, err := cloudevents.NewHTTPTransport(cloudevents.WithPort(port)); if err != nil {
		log.Fatalf("[CloudEvent] Failed to create HTTP transport layer, %v", err)
		return err
	}

	c, err := cloudevents.NewClient(t); if err != nil {
		log.Fatalf("[CloudEvent] Failed to create CloudEvent client, %v", err)
		return err
	}

	configureCodec()

	log.Printf("[CloudEvent] Listening for cloud events on port: %v...", port)

	return c.StartReceiver(context.Background(), unpackEvent(handler))
}
