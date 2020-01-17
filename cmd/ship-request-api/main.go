package main

import (
	"github.com/mhaddon/gke-knative/pkg/handler"
	"log"
)

func main() {
	log.Print("[Ship-Request-Api] Started.")

	config := getConfig()
	getPersistence().GetSession()

	log.Fatalf("[Ship-Request-Api][HTTP] %v",  handler.CreateHTTPListener(config.HttpPort, config.HttpOrigin, createRouter()))
}
