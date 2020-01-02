package normaliser

import (
	"fmt"
	"github.com/gorilla/mux"
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

	router.HandleFunc("/health", alwaysHealthy).Methods("GET")
	router.HandleFunc("/alive", alwaysHealthy).Methods("GET")

	return router
}

func alwaysHealthy(w http.ResponseWriter, r *http.Request) {
	printMessage(w, 200, []byte("1"))
}

func printMessage(w http.ResponseWriter, responseCode int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	if _, err := w.Write(response); err != nil {
		log.Fatalf("[Normaliser][HTTP] Error creating response: %v", err)
	}
}
