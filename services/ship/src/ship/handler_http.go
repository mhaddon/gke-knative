package ship

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateHTTPListener() error {
	config := getConfig()

	log.Printf("[Ship][HTTP] Listening for http requests on port: %v...", config.HTTP.Port)
	return http.ListenAndServe(fmt.Sprintf(":%v", config.HTTP.Port), createRouter())
}

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", alwaysHealthy).Methods("GET")
	router.HandleFunc("/alive", alwaysHealthy).Methods("GET")

	router.HandleFunc("/notifications", addNotification).Methods("PUT")
	router.HandleFunc("/notifications", getNotifications).Methods("GET")

	return router
}

func alwaysHealthy(w http.ResponseWriter, r *http.Request) {
	printMessage(w, 200, []byte("1"))
}
