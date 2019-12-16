// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package getbtc

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Get is the handler for GET /getbtc
// Send deposit address request
func (api GetbtcAPI) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var respBody []types.Bluewallet_types_BTCAddressItem
	json.NewEncoder(w).Encode(&respBody)
}
