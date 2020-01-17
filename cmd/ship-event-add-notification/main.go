package main

import (
	handler "github.com/mhaddon/gke-knative/pkg/handler"
	"log"
)

func main() {
	log.Print("[Ship-Event-Add-Notification] Started.")

	config := getConfig()
	getPersistence().GetSession()

	log.Fatalf("[Ship-Event-Add-Notification][CloudEvent] %v", handler.CreateCloudWatchListener(config.HttpPort, cloudWatchHandler))
}
