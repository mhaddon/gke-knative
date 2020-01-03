package ship

import (
	"encoding/json"
	"github.com/mhaddon/gke-knative/services/common/src/helper"
	"github.com/mhaddon/gke-knative/services/common/src/models"
	"log"
	"net/http"
)

func getNotifications(w http.ResponseWriter, r *http.Request) error {
	result := make([]models.ShipNotification, 0, 25)

	if err := getCollection().Find(nil).All(&result); err != nil {
		_ = helper.PrintErrorMessage(w, 500, "Could not process request")
		log.Printf("[Ship][DAO] Error getting ship notifications... %v", err)
		return err
	}

	data, err := json.Marshal(&result)

	if err != nil {
		_ = helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Print(err)
		return err
	}

	return helper.PrintMessage(w, 200, data)
}

func addNotification(w http.ResponseWriter, r *http.Request) error {
	shipNotification := models.ShipNotification{}

	if err := json.NewDecoder(r.Body).Decode(&shipNotification); err != nil {
		_ = helper.PrintErrorMessage(w, 400, "Invalid input body")
		log.Printf("[Ship][DAO] Error deserialising body... %v", err)
		return err
	}

	if err := getCollection().Insert(&shipNotification); err != nil {
		_ = helper.PrintErrorMessage(w, 400, "Failed to save data")
		log.Printf("[Ship][DAO] Error saving ship notification... %v", err)
		return err
	}

	data, err := json.Marshal(&shipNotification); if err != nil {
		_ = helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Printf("[Ship][DAO] Error serialising ship notification response... %v", err)
		return err
	}

	return helper.PrintMessage(w, 200, data)
}


