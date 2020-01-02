package main

import (
	"./normaliser"
	"log"
)

func main() {
	log.Print("[Normaliser] Started.")

	go func() { log.Fatalf("[Normaliser][CloudEvent] %v", normaliser.CreateCloudWatchListener()) }()
	go func() { log.Fatalf("[Normaliser][HTTP] %v",  normaliser.CreateHTTPListener()) }()
	select {}
}
