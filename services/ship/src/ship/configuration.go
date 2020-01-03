package ship

import (
	"encoding/json"
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"reflect"
	"sync"
)

type configuration struct {
	Mongo *mongoConfig `json:"mongo" file:"resources/persistence.json"`
	HTTP  *httpConfig `json:"http" file:"resources/http.json"`
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
}

var (
	configInstance *configuration
	configOnce     sync.Once
)

func test() *configuration {
	config := configuration{}

	t := reflect.TypeOf(config)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("file")

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		instance := reflect.New(reflect.TypeOf(field.Type))

		if err := processFile(tag, instance); err != nil {
			log.Printf("[Config] Failed to read %v config file at: %v, error: %v", field.Name, tag, err)
		}

		reflect.ValueOf(config).Elem().FieldByName(field.Name).Set(instance)
	}

	return &config
}

func (c *configuration) convertToJson() string {
	b, err := json.Marshal(c); if err != nil {
		log.Printf("[Ship][Config] Failed to stringify configuration: %v", err)
	}

	return string(b)
}

func getConfig() *configuration {
	configOnce.Do(func() {
		configInstance = test()

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
