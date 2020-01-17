package main

import (
	"errors"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/mhaddon/gke-knative/services/common/pkg/models"
	"log"
	"strconv"
	"strings"
)

type shipNotificationSpeedVersionOne struct {
	Velocity float64 `json:"velocity"`
	Unit     string  `json:"unit"`
}

type shipNotificationVersionOne struct {
	StandardVersion int    `json:"version"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Captain         string `json:"captain"`
	Position        string `json:"position"`

	Speed shipNotificationSpeedVersionOne `json:"speed"`
}

type versionOne struct{}

func convertVelocityToKnots(unit string, velocity float64) (*float64, error) {
	velocityInKnots := 0.0

	switch u := unit; u {
	case "Knots":
		velocityInKnots = velocity
	case "Miles":
		velocityInKnots = velocity * 0.868976
	default:
		return nil, errors.New("invalid velocity unit")
	}

	return &velocityInKnots, nil
}

func convertCoordinatesToPosition(position string) (*float64, *float64, error) {
	coords := strings.Split(position, ",")
	lat, err := strconv.ParseFloat(coords[0], 64); if err != nil {
		return nil, nil, err
	}
	long, err := strconv.ParseFloat(coords[0], 64); if err != nil {
		return nil, nil, err
	}

	return &lat, &long, nil
}

func convertVersionOneToNotification(payload shipNotificationVersionOne) (*models.ShipNotification, error) {
	lat, long, err := convertCoordinatesToPosition(payload.Position); if err != nil {
		log.Printf("[Normaliser][VersionOne] Error Converting Position: %s\n", err.Error())
		return nil, err
	}

	velocity, err := convertVelocityToKnots(payload.Speed.Unit, payload.Speed.Velocity); if err != nil {
		log.Printf("[Normaliser][VersionOne] Error Converting Velocity: %s\n", err.Error())
		return nil, err
	}

	result := models.ShipNotification{
		Registration: models.ShipNotificationRegistration{
			Name:    payload.Name,
			Captain: payload.Captain,
			ID:      payload.ID,
		},
		Status: models.ShipNotificationStatus{
			Position: models.ShipNotificationPosition{
				Lat:  *lat,
				Long: *long,
			},
			Velocity: *velocity,
		},
	}

	return &result, nil
}

func (versionOne) apply(event cloudevents.Event) (*models.ShipNotification, error) {
	payload := &shipNotificationVersionOne{}

	if err := event.DataAs(payload); err != nil {
		log.Printf("[Normaliser][VersionOne] Error deserialising event: %s\n", err.Error())
		return nil, err
	}

	result, err := convertVersionOneToNotification(*payload); if err != nil {
		log.Printf("[Normaliser][VersionOne] Error normalising data: %s\n", err.Error())
		return nil, err
	}

	return result, nil
}
