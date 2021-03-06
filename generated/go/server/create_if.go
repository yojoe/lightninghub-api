// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// CreateInterface is interface for /create root endpoint
type CreateInterface interface { // Post is the handler for POST /create
	// Send account creation request
	Post(http.ResponseWriter, *http.Request)
}

// CreateInterfaceRoutes is routing for /create root endpoint
func CreateInterfaceRoutes(r *mux.Router, i CreateInterface) {
	r.HandleFunc("/create", i.Post).Methods("POST")
}
