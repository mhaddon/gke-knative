package main

import (
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/mhaddon/gke-knative/pkg/models"
)

type operator interface {
	apply(cloudevents.Event) (*models.ShipNotification, error)
}

type operation struct {
	Operator operator
}

func (o *operation) operate(event cloudevents.Event) (*models.ShipNotification, error) {
	return o.Operator.apply(event)
}
