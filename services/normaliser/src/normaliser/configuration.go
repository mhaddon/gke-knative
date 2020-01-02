package normaliser

import (
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"sync"
)

type configuration struct {
	Egress     *egress
	HTTP       *httpConfig
	CloudEvent *cloudEventConfig
}

type egress struct {
	Source string `env:"Source"`
	Type   string `env:"Type"`
}

type httpConfig struct {
	Port int `env:"HTTP_PORT"`
}

type cloudEventConfig struct {
	Port int `env:"HTTP_PORT"`
}

var (
	configInstance *configuration
	configOnce     sync.Once
)

func getConfig() *configuration {
	configOnce.Do(func() {
		configInstance = &configuration{
			Egress:     createEgressConfig(),
			HTTP:       createHTTPConfig(),
			CloudEvent: createCloudEventConfig(),
		}
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

func createEgressConfig() *egress {
	egress := &egress{}

	if err := processFile("resources/egress.json", egress); err != nil {
		log.Printf("[Normaliser][Config] Failed to read egress config file: %v", err)
	}

	return egress
}

func createHTTPConfig() *httpConfig {
	http := &httpConfig{}

	if err := processFile("resources/http.json", http); err != nil {
		log.Printf("[Normaliser][Config] Failed to read http config file: %v", err)
	}

	return http
}

func createCloudEventConfig() *cloudEventConfig {
	cloudevent := &cloudEventConfig{}

	if err := processFile("resources/cloudevent.json", cloudevent); err != nil {
		log.Printf("[Normaliser][Config] Failed to read cloudevent config file: %v", err)
	}

	return cloudevent
}
