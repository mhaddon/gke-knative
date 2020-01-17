package main

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/mhaddon/gke-knative/pkg/persistence"
	"log"
	"sync"
)

type configuration struct {
	HttpPort int `env:"PORT" envDefault:"8080" json:"http_port"`

	MongoDomain     string `env:"MONGO_DOMAIN" envDefault:"localhost" json:"mongo_domain"`
	MongoPort       int    `env:"MONGO_PORT" envDefault:"27017" json:"mongo_port"`
	MongoDB         string `env:"MONGO_DB" envDefault:"ship" json:"mongo_db"`
	MongoCollection string `env:"MONGO_COLLECTION" envDefault:"app" json:"mongo_collection"`
	MongoUsername   string `env:"MONGO_USERNAME" envDefault:"root" json:"mongo_username"`
	MongoPassword   string `env:"MONGO_PASSWORD" envDefault:"password" json:"-"`
}

var (
	configInstance *configuration
	configOnce     sync.Once

	persistenceInstance *persistence.Mongo
	persistenceOnce     sync.Once
)

func (c *configuration) convertToJson() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Printf("[Ship-Event-Add-Notification][Config] Failed to stringify configuration: %v", err)
	}

	return string(b)
}

func getConfig() *configuration {
	configOnce.Do(func() {
		if err := env.Parse(&configInstance); err != nil {
			fmt.Printf("[Ship-Event-Add-Notification][Config] Failed to process environments: %v", err)
		}

		log.Printf("[Ship-Event-Add-Notification][Config] Loaded config with variables: %v", configInstance.convertToJson())
	})
	return configInstance
}

func getPersistence() *persistence.Mongo {
	config := getConfig()

	persistenceOnce.Do(func() {
		persistenceInstance = &persistence.Mongo{
			MongoDomain:     config.MongoDomain,
			MongoPort:       config.MongoPort,
			MongoDB:         config.MongoDB,
			MongoCollection: config.MongoCollection,
			MongoUsername:   config.MongoUsername,
			MongoPassword:   config.MongoPassword,
		}
	})

	return persistenceInstance
}
