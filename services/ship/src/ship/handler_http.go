package ship

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-knative/services/common/src/helper"
	"log"
	"net/http"
)

func CreateHTTPListener() error {
	config := getConfig()

	log.Printf("[Ship][HTTP] Listening for http requests on port: %v...", config.HTTP.Port)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{config.HTTP.Origin})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	handler := handlers.CORS(originsOk, headersOk, methodsOk)

	return http.ListenAndServe(fmt.Sprintf(":%v", config.HTTP.Port), handler(createRouter()))
}

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", helper.LogQuery(helper.AlwaysHealthy, "[Ship]")).Methods("GET")
	router.HandleFunc("/alive", helper.LogQuery(helper.AlwaysHealthy, "[Ship]")).Methods("GET")

	router.HandleFunc("/notifications", helper.LogQuery(addNotification, "[Ship]")).Methods("PUT")
	router.HandleFunc("/notifications", helper.LogQuery(getNotifications, "[Ship]")).Methods("GET")

	return router
}
