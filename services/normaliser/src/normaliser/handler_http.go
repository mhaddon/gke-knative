package normaliser

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-knative/services/common/src/helper"
	"log"
	"net/http"
)

func CreateHTTPListener() error {
	config := getConfig()

	log.Printf("[Normaliser][HTTP] Listening for http requests on port: %v...", config.HTTP.Port)
	return http.ListenAndServe(fmt.Sprintf(":%v", config.HTTP.Port), createRouter())
}

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", helper.LogQuery(helper.AlwaysHealthy, "[Normaliser]")).Methods("GET")
	router.HandleFunc("/alive", helper.LogQuery(helper.AlwaysHealthy, "[Normaliser]")).Methods("GET")

	return router
}
