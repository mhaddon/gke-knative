package main

import (
	connector "./ship_notification_connector"
	"log"
)

func main() {
	log.Print("[ShipNotificationConnector] Started.")

	log.Fatalf("[ShipNotificationConnector][CloudEvent] %v",  connector.CreateCloudWatchListener())
}
