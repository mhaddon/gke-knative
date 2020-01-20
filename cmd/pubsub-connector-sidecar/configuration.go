package main

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env"
	"log"
	"sync"
)

type configuration struct {
	TargetUrl string `env:"TARGET_URL" envDefault:"http://localhost:80/" json:"target_url"`

	SubscriptionId   string `env:"SUBSCRIPTION_ID" envDefault:"default" json:"subscription_id"`
	TopicName        string `env:"TOPIC_NAME" envDefault:"default" json:"topic_name"`
	IngressTopicName string `env:"INGRESS_TOPIC_NAME" envDefault:"default" json:"ingress_topic_name"`
	EgressTopicName  string `env:"EGRESS_TOPIC_NAME" envDefault:"default2" json:"egress_topic_name"`
	ProjectId        string `env:"PROJECT_ID" envDefault:"default" json:"project_name"`
}

var (
	configInstance *configuration
	configOnce     sync.Once
)

func (c *configuration) convertToJson() string {
	b, err := json.Marshal(c); if err != nil {
		log.Printf("[PubSub-Connector-Sidecar][Config] Failed to stringify configuration: %v", err)
	}

	return string(b)
}

func getConfig() *configuration {
	configOnce.Do(func() {
		config := configuration{}
		if err := env.Parse(&config); err != nil {
			fmt.Printf("[PubSub-Connector-Sidecar][Config] Failed to process environments: %v", err)
		}
		configInstance = &config

		log.Printf("[PubSub-Connector-Sidecar][Config] Loaded config with variables: %v", configInstance.convertToJson())
	})
	return configInstance
}
