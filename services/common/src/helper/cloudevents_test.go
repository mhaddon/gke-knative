package helper

import (
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/mhaddon/gke-knative/services/common/src/models"
	"testing"
)

func unpackTest(inputJson string, t *testing.T) {
	event := cloudevents.NewEvent()

	if err := event.SetData(inputJson); err != nil {
		t.Error("Error initialising test")
	}

	eventResponse, err := UnpackNestedCloudEvent(&event); if err != nil {
		t.Error("Error unpacking event")
	}

	shipNotification := &models.ShipNotification{}
	if err := eventResponse.DataAs(shipNotification); err != nil {
		t.Errorf("Error converting data to ship notification")
	}

	// todo - more extensive testing
	if shipNotification.Registration.Name != "Golden Sausage" {
		t.Error("Failed to convert name")
	}
}

func TestUnpackNestedCloudEventString(t *testing.T) {
	eventDataJson := "{\"cloudEventsVersion\":\"0.1\",\"contentType\":\"application/json\",\"data\":\"{\\\"registration\\\":{\\\"id\\\":\\\"1x0d\\\",\\\"name\\\":\\\"Golden Sausage\\\",\\\"captain\\\":\\\"Duke Lordship\\\"},\\\"status\\\":{\\\"position\\\":{\\\"lat\\\":49.2144,\\\"long\\\":49.2144},\\\"velocity\\\":40}}\",\"eventID\":\"ccf77693-ae0f-494a-a019-5e7c799abea5\",\"eventTime\":\"2020-01-05T18:12:58.982994654Z\",\"eventType\":\"shipdata/normalised\",\"extensions\":{\"Knativearrivaltime\":\"2020-01-05T18:12:58Z\",\"Knativebrokerttl\":254,\"Knativehistory\":\"default-kne-ingress-kn-channel.default.svc.cluster.local\",\"Traceparent\":\"00-d114051e5bfede9feabe7297eaf5da0d-de45910a4ed68544-00\",\"traceparent\":\"00-d114051e5bfede9feabe7297eaf5da0d-d4d3e7556f79cbde-00\"},\"source\":\"mhaddon/normaliser\"}"

	unpackTest(eventDataJson, t)
}

func TestUnpackNestedCloudEvent(t *testing.T) {
	eventDataJson := "{\"cloudEventsVersion\":\"0.1\",\"contentType\":\"application/json\",\"data\":{\"registration\":{\"id\":\"1x0d\",\"name\":\"Golden Sausage\",\"captain\":\"Duke Lordship\"},\"status\":{\"position\":{\"lat\":49.2144,\"long\":49.2144},\"velocity\":40}},\"eventID\":\"ccf77693-ae0f-494a-a019-5e7c799abea5\",\"eventTime\":\"2020-01-05T18:12:58.982994654Z\",\"eventType\":\"shipdata/normalised\",\"extensions\":{\"Knativearrivaltime\":\"2020-01-05T18:12:58Z\",\"Knativebrokerttl\":254,\"Knativehistory\":\"default-kne-ingress-kn-channel.default.svc.cluster.local\",\"Traceparent\":\"00-d114051e5bfede9feabe7297eaf5da0d-de45910a4ed68544-00\",\"traceparent\":\"00-d114051e5bfede9feabe7297eaf5da0d-d4d3e7556f79cbde-00\"},\"source\":\"mhaddon/normaliser\"}"

	unpackTest(eventDataJson, t)
}

func TestUnpackUnnestedCloudEvent(t *testing.T) {
	eventDataJson := "{ \"registration\": { \"id\": \"1x0d\", \"name\": \"Golden Sausage\", \"captain\": \"Duke Lordship\" }, \"status\": { \"position\": { \"lat\":49.2144, \"long\":2.1312 }, \"velocity\":41 } }"

	unpackTest(eventDataJson, t)
}