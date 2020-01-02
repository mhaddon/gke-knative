package ship_notification_connector

import (
	"encoding/json"
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"sync"
)

type configuration struct {
	Services   *servicesConfig
	CloudEvent *cloudEventConfig `json:"cloudevent"`
}

type servicesConfig struct {
	Ship string `env:"SHIP_DOMAIN" json:"ship"`
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
		log.Printf("[ShipNotificationConnector][Config] Failed to stringify configuration: %v", err)
	}

	return string(b)
}

func getConfig() *configuration {
	configOnce.Do(func() {
		configInstance = &configuration{
			Services:   createServicesConfig(),
			CloudEvent: createCloudEventConfig(),
		}

		log.Printf("[ShipNoitifcationConnector][Config] Loaded config with variables: %v", configInstance.convertToJson())
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

func createServicesConfig() *servicesConfig {
	servicesConfig := &servicesConfig{}

	if err := processFile("resources/persistence.json", servicesConfig); err != nil {
		log.Printf("[ShipNoitifcationConnector][Config] Failed to read services config file: %v", err)
	}

	return servicesConfig
}

func createCloudEventConfig() *cloudEventConfig {
	cloudEventConfig := &cloudEventConfig{}

	if err := processFile("resources/cloudevent.json", cloudEventConfig); err != nil {
		log.Printf("[ShipNoitifcationConnector][Config] Failed to read cloudevent config file: %v", err)
	}

	return cloudEventConfig
}
