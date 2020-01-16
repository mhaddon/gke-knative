package main

import (
	handler "github.com/mhaddon/gke-knative/services/common/pkg/handler/cloudevents"
	"log"
)

func main() {
	log.Print("[Normaliser-Event-Normaliser] Started.")

	config := getConfig()
	log.Fatalf("[Normaliser-Event-Normaliser][CloudEvent] %v", handler.CreateCloudWatchListener(config.HttpPort, cloudWatchHandler))
}
