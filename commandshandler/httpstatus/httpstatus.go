package httpstatus

import "net/http"

// SendBadrequest responds to the request by sending a bad request error
func SendBadrequest(w http.ResponseWriter, reason string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - " + reason))
}

// SendInternalServerError responds to the request by sending an internal server error
func SendInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("500 - An internal server error occured"))
}
