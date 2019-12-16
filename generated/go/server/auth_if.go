// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// AuthInterface is interface for /auth root endpoint
type AuthInterface interface { // Post is the handler for POST /auth
	// Send user authentication request
	Post(http.ResponseWriter, *http.Request)
}

// AuthInterfaceRoutes is routing for /auth root endpoint
func AuthInterfaceRoutes(r *mux.Router, i AuthInterface) {
	r.HandleFunc("/auth", i.Post).Methods("POST")
}