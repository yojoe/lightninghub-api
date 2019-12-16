// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// GetinfoInterface is interface for /getinfo root endpoint
type GetinfoInterface interface { // Get is the handler for GET /getinfo
	// Send status information request
	Get(http.ResponseWriter, *http.Request)
}

// GetinfoInterfaceRoutes is routing for /getinfo root endpoint
func GetinfoInterfaceRoutes(r *mux.Router, i GetinfoInterface) {
	r.HandleFunc("/getinfo", i.Get).Methods("GET")
}