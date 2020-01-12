package ship

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-knative/services/common/src/helper"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func CreateHTTPListener() error {
	config := getConfig()

	log.Printf("[Ship][HTTP] Listening for http requests on port: %v...", config.HTTP.Port)
	return http.ListenAndServe(fmt.Sprintf(":%v", config.HTTP.Port), enableCors(createRouter()))
}

func enableCors(router *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //todo change
		AllowedMethods: []string{"GET", "POST", "PUT"},
		Debug: true, // todo change for prod
	})

	return c.Handler(router)
}

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", helper.LogQuery(helper.AlwaysHealthy, "[Ship]")).Methods("GET")
	router.HandleFunc("/alive", helper.LogQuery(helper.AlwaysHealthy, "[Ship]")).Methods("GET")

	router.HandleFunc("/notifications", helper.LogQuery(addNotification, "[Ship]")).Methods("PUT")
	router.HandleFunc("/notifications", helper.LogQuery(getNotifications, "[Ship]")).Methods("GET")

	return router
}
