package ship

import (
	"encoding/json"
	"github.com/mhaddon/gke-knative/services/common/src/models"
	"log"
	"net/http"
)

func getNotifications(w http.ResponseWriter, r *http.Request) {
	result := make([]models.ShipNotification, 0, 25)

	if err := getCollection().Find(nil).All(&result); err != nil {
		printErrorMessage(w, 500, "Could not process request")
		log.Printf("[Ship][DAO] Error getting ship notifications... %v", err)
		return
	}

	data, err := json.Marshal(&result)

	if err != nil {
		printErrorMessage(w, 500,"Could not process response")
		log.Print(err)
		return
	}

	printMessage(w, 200, data)
}

func addNotification(w http.ResponseWriter, r *http.Request) {
	shipNotification := models.ShipNotification{}

	if err := json.NewDecoder(r.Body).Decode(&shipNotification); err != nil {
		printErrorMessage(w, 400, "Invalid input body")
		log.Printf("[Ship][DAO] Error deserialising body... %v", err)
		return
	}

	if err := getCollection().Insert(&shipNotification); err != nil {
		printErrorMessage(w, 400, "Failed to save data")
		log.Printf("[Ship][DAO] Error saving ship notification... %v", err)
		return
	}

	data, err := json.Marshal(&shipNotification); if err != nil {
		printErrorMessage(w, 500,"Could not process response")
		log.Printf("[Ship][DAO] Error serialising ship notification response... %v", err)
		return
	}

	printMessage(w, 200, data)
}

func printErrorMessage(w http.ResponseWriter, responseCode int, response string) {
	message := map[string]interface{}{ "err": response, "code": responseCode }
	encodedMessage, _ := json.Marshal(message)
	printMessage(w, responseCode, encodedMessage)
}

func printMessage(w http.ResponseWriter, responseCode int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	if _, err := w.Write(response); err != nil {
		log.Fatalf("[Ship][HTTP] Error creating response: %v", err)
	}
}
