package normaliser

import (
	"encoding/json"
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"sync"
)

type configuration struct {
	Egress     *egress `json:"egress"`
	HTTP       *httpConfig `json:"http"`
	CloudEvent *cloudEventConfig `json:"cloudevent"`
}

type egress struct {
	Source string `env:"EGRESS_SOURCE" json:"source"`
	Type   string `env:"EGRESS_TYPE" json:"type"`
}

type httpConfig struct {
	Port int `env:"HTTP_PORT" json:"port"`
}

type cloudEventConfig struct {
	Port int `env:"CLOUDEVENT_PORT" json:"port"`
}

var (
	configInstance *configuration
	configOnce     sync.Once
)

func (c *configuration) convertToJson() string {
	b, err := json.Marshal(c); if err != nil {
		log.Printf("[Normaliser][Config] Failed to stringify configuration: %v", err)
	}

	return string(b)
}

func getConfig() *configuration {
	configOnce.Do(func() {
		configInstance = &configuration{
			Egress:     createEgressConfig(),
			HTTP:       createHTTPConfig(),
			CloudEvent: createCloudEventConfig(),
		}

		log.Printf("[Normaliser][Config] Loaded config with variables: %v", configInstance.convertToJson())
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
