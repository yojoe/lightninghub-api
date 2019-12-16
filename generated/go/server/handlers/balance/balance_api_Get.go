// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package balance

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Get is the handler for GET /balance
// Send balance information request
func (api BalanceAPI) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var respBody types.BalanceGetRespBody
	json.NewEncoder(w).Encode(&respBody)
}
