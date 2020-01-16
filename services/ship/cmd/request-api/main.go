package main

import (
	handler "github.com/mhaddon/gke-knative/services/common/pkg/handler/http"
	"log"
)

func main() {
	log.Print("[Ship-Request-Api] Started.")

	config := getConfig()
	getPersistence().GetSession()

	log.Fatalf("[Ship-Request-Api][HTTP] %v",  handler.CreateHTTPListener(config.HttpPort, config.HttpOrigin, createRouter()))
}
