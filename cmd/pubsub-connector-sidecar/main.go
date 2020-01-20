package main

import (
	"log"
)

func main()  {
	log.Print("[PubSub-Connector-Sidecar] Started.")

	// todo this code-base needs to be restructured as the flow is too cognitively complex
	log.Fatalf("[PubSub-Connector-Sidecar][PubSub] %v", pull(sendCloudEvent))
}
