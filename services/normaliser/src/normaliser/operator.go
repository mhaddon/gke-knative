package normaliser

import (
	models "github.com/mhaddon/gke-knative/services/common/src/models"
	cloudevents "github.com/cloudevents/sdk-go"
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
