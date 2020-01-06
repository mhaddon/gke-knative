package ship_notification_connector

import (
	"log"
	"testing"
	"github.com/mhaddon/gke-knative/services/common/src/models"
	"io/ioutil"
	"fmt"
)

func TestCreatePutRequest(t *testing.T) {
	configOnce.Do(func() {
		configInstance = &configuration{
			Services:   &servicesConfig {
				Ship: "http://localhost",
			},
		}
	})

	notification := models.ShipNotification {
		Registration: models.ShipNotificationRegistration {
			Name: "TestName",
			Captain: "TestCaptain",
			ID: "TestID",
		},
		Status: models.ShipNotificationStatus{
			Position: models.ShipNotificationPosition{
				Lat:  1.0,
				Long: 2.0,
			},
			Velocity: 10,
		},
	}

	// todo: there is not actually any tests here...

	request, err := createPutRequest(&notification); if err != nil {
		log.Printf("err: %v", err)
	}

	contents, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The calculated length is:", len(string(contents)))
}