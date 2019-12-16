// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package gettxs

import (
	"encoding/json"
	"net/http"
)

// Get is the handler for GET /gettxs
// Get a list of transactions (deposits and payments) along with their details
// for the authenticated user's custody account.
func (api GettxsAPI) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var respBody []json.RawMessage
	json.NewEncoder(w).Encode(&respBody)
}
