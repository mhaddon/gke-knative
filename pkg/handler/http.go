package handler

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateHTTPListener(port int, origin string, router *mux.Router) error {
	log.Printf("[HTTP] Listening for http requests on port: %v...", port)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{origin})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	handler := handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())

	return http.ListenAndServe(fmt.Sprintf(":%v", port), handler(router))
}
