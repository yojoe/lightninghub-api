// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package getuserinvoices

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Get is the handler for GET /getuserinvoices
// Get a list of Lightning invoices along with their payment status details
// for the authenticated user's custody account.
func (api GetuserinvoicesAPI) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var respBody []types.Bluewallet_types_Invoice
	json.NewEncoder(w).Encode(&respBody)
}
