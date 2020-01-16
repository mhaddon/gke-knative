package main

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env"
	"log"
	"sync"
)

type configuration struct {
	HttpPort     int    `env:"PORT" envDefault:"8080" json:"http_port"`
	EgressSource string `env:"EGRESS_SOURCE" envDefault:"mhaddon/normaliser" json:"egress_source"`
	EgressType   string `env:"EGRESS_TYPE" envDefault:"shipdata/normalised" json:"egress_type"`
}

var (
	configInstance *configuration
	configOnce     sync.Once
)

func (c *configuration) convertToJson() string {
	b, err := json.Marshal(c); if err != nil {
		log.Printf("[Normaliser-Event-Normaliser][Config] Failed to stringify configuration: %v", err)
	}

	return string(b)
}

func getConfig() *configuration {
	configOnce.Do(func() {
		if err := env.Parse(&configInstance); err != nil {
			fmt.Printf("[Normaliser-Event-Normaliser][Config] Failed to process environments: %v", err)
		}

		log.Printf("[Normaliser-Event-Normaliser][Config] Loaded config with variables: %v", configInstance.convertToJson())
	})
	return configInstance
}
