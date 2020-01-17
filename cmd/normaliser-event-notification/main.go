package main

import (
	"github.com/mhaddon/gke-knative/pkg/handler"
	"log"
)

func main() {
	log.Print("[Normaliser-Event-Normaliser] Started.")

	config := getConfig()
	log.Fatalf("[Normaliser-Event-Normaliser][CloudEvent] %v", handler.CreateCloudWatchListener(config.HttpPort, cloudWatchHandler))
}
