// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// DecodeinvoiceInterface is interface for /decodeinvoice root endpoint
type DecodeinvoiceInterface interface { // Get is the handler for GET /decodeinvoice
	// Decode a BOLT-11 compatible Lightning invoice into a human readable format.
	// Uses the implementation of the connected Lightning node so the client app
	// doesn't have to implement this.
	Get(http.ResponseWriter, *http.Request)
}

// DecodeinvoiceInterfaceRoutes is routing for /decodeinvoice root endpoint
func DecodeinvoiceInterfaceRoutes(r *mux.Router, i DecodeinvoiceInterface) {
	r.HandleFunc("/decodeinvoice", i.Get).Methods("GET")
}