package models

import (
	"encoding/json"
)

type ShipNotificationRegistration struct {
	ID      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Captain string `json:"captain" bson:"captain"`
}

type ShipNotificationPosition struct {
	Lat  float64 `json:"lat" bson:"lat"`
	Long float64 `json:"long" bson:"long"`
}

type ShipNotificationStatus struct {
	Position ShipNotificationPosition `json:"position" bson:"position"`
	Velocity float64                  `json:"velocity" bson:"velocity"`
}

type ShipNotification struct {
	Registration ShipNotificationRegistration `json:"registration" bson:"registration"`
	Status       ShipNotificationStatus       `json:"status" bson:"status"`
}

func (c *ShipNotification) ConvertToJson() (string, error) {
	b, err := json.Marshal(c); if err != nil {
		return "[]", err
	}

	return string(b), nil
}
