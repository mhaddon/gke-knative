package main

import (
	"github.com/mhaddon/gke-knative/pkg/handler"
	"log"
)

func main() {
	log.Print("[Normaliser-Event-Notification] Started.")

	config := getConfig()
	log.Fatalf("[Normaliser-Event-Notification][CloudEvent] %v", handler.CreateCloudWatchListener(config.HttpPort, cloudWatchHandler))
}
