package normaliser

import (
	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/mhaddon/gke-knative/services/common/src/models"
	"github.com/pkg/errors"
	"log"
	"sync"
)

type normalisers map[int]operation

type shipNotificationMetaData struct {
	StandardVersion int `json:"version"`
}

var (
	normaliserOnce sync.Once
	instance       normalisers
)

func getNormalisers() normalisers {
	normaliserOnce.Do(func() {
		instance = map[int]operation{
			1: {Operator: versionOne{}},
			2: {Operator: versionTwo{}},
		}
	})

	return instance
}

func getNormaliser(version int) (*operation, error) {
	normalisers := getNormalisers()

	normaliser, ok := normalisers[version]; if ok == false {
		log.Printf("[Normaliser][getNormaliser] Normaliser with this version does not exist. Version: %v\n", version)
		return nil, errors.New("unknown normaliser")
	}

	return &normaliser, nil
}

func getVersion(event cloudevents.Event) (int, error) {
	metadata := &shipNotificationMetaData{}
	if err := event.DataAs(metadata); err != nil {
		log.Printf("[Normaliser][getVersion] Error deserialising event: %s\n", err.Error())
		return 0, err
	}

	return metadata.StandardVersion, nil
}

func normaliseEvent(event cloudevents.Event) (*models.ShipNotification, error) {
	version, err := getVersion(event); if err != nil {
		log.Printf("[Normaliser][NormaliseEvent] Error getting event version: %s\n", err.Error())
		return nil, err
	}

	normaliser, err := getNormaliser(version); if err != nil {
		log.Printf("[Normaliser][NormaliseEvent] Error getting event normaliser: %s\n", err.Error())
		return nil, err
	}

	notification, err := normaliser.Operator.apply(event); if err != nil {
		log.Printf("[Normaliser][NormaliseEvent] Error normalising event: %s\n", err.Error())
		return nil, err
	}

	return notification, nil
}
