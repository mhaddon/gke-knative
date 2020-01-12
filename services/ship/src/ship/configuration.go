package ship

import (
	"encoding/json"
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"sync"
)

type configuration struct {
	Mongo *mongoConfig `json:"mongo"`
	HTTP  *httpConfig `json:"http"`
}

type mongoConfig struct {
	Domain     string `env:"MONGO_DOMAIN" json:"domain"`
	Port       int    `env:"MONGO_PORT" json:"port"`
	DB         string `env:"MONGO_DB" json:"db"`
	Collection string `env:"MONGO_COLLECTION" json:"collection"`
	Username   string `env:"MONGO_USERNAME" json:"username"`
	Password   string `env:"MONGO_PASSWORD" json:"-"`
}

type httpConfig struct {
	Port int `env:"HTTP_PORT" json:"port"`
	Origin string `env:"CORS_ORIGIN" json:"origin"`
}

var (
	configInstance *configuration
	configOnce     sync.Once
)

func (c *configuration) convertToJson() string {
	b, err := json.Marshal(c); if err != nil {
		log.Printf("[Ship][Config] Failed to stringify configuration: %v", err)
	}

	return string(b)
}

func getConfig() *configuration {
	configOnce.Do(func() {
		configInstance = &configuration{
			Mongo: createMongoConfig(),
			HTTP:  createHTTPConfig(),
		}

		log.Printf("[Ship][Config] Loaded config with variables: %v", configInstance.convertToJson())
	})
	return configInstance
}

func processFile(path string, data interface{}) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if err := gonfig.GetConf(path, data); err != nil {
			return err
		}
	}

	return nil
}

func createMongoConfig() *mongoConfig {
	mongoConfig := &mongoConfig{}

	if err := processFile("resources/persistence.json", mongoConfig); err != nil {
		log.Printf("[Ship][Config] Failed to read mongo config file: %v", err)
	}

	return mongoConfig
}

func createHTTPConfig() *httpConfig {
	http := &httpConfig{}

	if err := processFile("resources/http.json", http); err != nil {
		log.Printf("[Ship][Config] Failed to read http config file: %v", err)
	}

	return http
}
