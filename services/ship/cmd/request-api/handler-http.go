package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-knative/services/common/pkg/helper"
	"github.com/mhaddon/gke-knative/services/common/pkg/models"
	"log"
	"net/http"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", helper.LogQuery(helper.AlwaysHealthy, "[Ship-Request-Api]")).Methods("GET")
	router.HandleFunc("/alive", helper.LogQuery(helper.AlwaysHealthy, "[Ship-Request-Api]")).Methods("GET")

	router.HandleFunc("/notifications", helper.LogQuery(getNotifications, "[Ship-Request-Api]")).Methods("GET")

	return router
}

func getNotifications(w http.ResponseWriter, r *http.Request) error {
	result := make([]models.ShipNotification, 0, 25)

	if err := getPersistence().GetCollection().Find(nil).All(&result); err != nil {
		_ = helper.PrintErrorMessage(w, 500, "Could not process request")
		log.Printf("[Ship-Request-Api][DAO] Error getting ship notifications... %v", err)
		return err
	}

	data, err := json.Marshal(&result); if err != nil {
		_ = helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Printf("[Ship-Request-Api][DAO] Error getting ship notifications... %v", err)
		return err
	}

	return helper.PrintMessage(w, 200, data)
}