package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func LogQuery(callback func(w http.ResponseWriter, r *http.Request) error, logPrefix string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v[HTTP] Received a %v request for: %v, from: %v, with size: %v...", logPrefix, r.Method, r.RequestURI, r.RemoteAddr, r.ContentLength)

		startNanoSeconds := time.Now().UnixNano()

		if err := callback(w, r); err != nil {
			log.Printf("%v[HTTP] %v Request for: %v, from %v, with size: %v has resulted in an error: %v", logPrefix, r.Method, r.RequestURI, r.RemoteAddr, r.ContentLength, err)
		}

		durationNanoseconds := time.Now().UnixNano() - startNanoSeconds

		log.Printf("%v[HTTP] %v Request for: %v, from: %v, with size: %v took %v nanoseconds", logPrefix, r.Method, r.RequestURI, r.RemoteAddr, r.ContentLength, durationNanoseconds)
	}
}

func PrintErrorMessage(w http.ResponseWriter, responseCode int, response string) error {
	message := map[string]interface{}{ "err": response, "code": responseCode }
	encodedMessage, _ := json.Marshal(message)
	return PrintMessage(w, responseCode, encodedMessage)
}

func PrintMessage(w http.ResponseWriter, responseCode int, response []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	if _, err := w.Write(response); err != nil {
		log.Printf("[PrintMessage] Error creating response: %v", err)
		return err
	}

	return nil
}

func AlwaysHealthy(w http.ResponseWriter, r *http.Request) error {
	return PrintMessage(w, 200, []byte("1"))
}
