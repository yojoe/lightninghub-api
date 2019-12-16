// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package getpending

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Get is the handler for GET /getpending
// Get a list of pending deposits along with their details
// for the authenticated user's custody account.
func (api GetpendingAPI) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var respBody []types.Bluewallet_types_BTCTransaction
	json.NewEncoder(w).Encode(&respBody)
}
