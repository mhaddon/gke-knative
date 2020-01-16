package main

import (
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/mhaddon/gke-knative/services/common/pkg/models"
	"log"
)

type shipNotificationPositionVersionTwo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type shipNotificationShipVersionTwo struct {
	Registration string `json:"registration"`
	Name         string `json:"name"`
}

type shipNotificationVersionTwo struct {
	Ship     shipNotificationShipVersionTwo     `json:"ship"`
	Captain  string                             `json:"captain"`
	Position shipNotificationPositionVersionTwo `json:"position"`
	Speed    float64                            `json:"speed"`
}

type versionTwo struct{}

func convertVersionTwoToNotification(payload shipNotificationVersionTwo) (*models.ShipNotification, error) {
	result := models.ShipNotification{
		Registration: models.ShipNotificationRegistration{
			Name:    payload.Ship.Name,
			Captain: payload.Captain,
			ID:      payload.Ship.Registration,
		},
		Status: models.ShipNotificationStatus{
			Position: models.ShipNotificationPosition{
				Lat:  payload.Position.Latitude,
				Long: payload.Position.Longitude,
			},
			Velocity: payload.Speed,
		},
	}

	return &result, nil
}

func (versionTwo) apply(event cloudevents.Event) (*models.ShipNotification, error) {
	payload := &shipNotificationVersionTwo{}

	if err := event.DataAs(payload); err != nil {
		log.Printf("[Normaliser][VersionTwo] Error deserialising event: %s\n", err.Error())
		return nil, err
	}

	result, err := convertVersionTwoToNotification(*payload); if err != nil {
		log.Printf("[Normaliser][VersionTwo] Error normalising data: %s\n", err.Error())
		return nil, err
	}

	return result, nil
}
