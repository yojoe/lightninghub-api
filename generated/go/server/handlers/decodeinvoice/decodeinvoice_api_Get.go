// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package decodeinvoice

import (
	"encoding/json"
	bluewallet_types_types "github.com/yojoe/lightninghub-api/go/server/types/bluewallet_types/types"
	"net/http"
)

// Get is the handler for GET /decodeinvoice
// Decode a BOLT-11 compatible Lightning invoice into a human readable format.
// Uses the implementation of the connected Lightning node so the client app
// doesn't have to implement this.
func (api DecodeinvoiceAPI) Get(w http.ResponseWriter, r *http.Request) { // invoice := req.FormValue("invoice")
	w.Header().Set("Content-Type", "application/json")
	var respBody bluewallet_types_types.LNDPayReq
	json.NewEncoder(w).Encode(&respBody)
}
