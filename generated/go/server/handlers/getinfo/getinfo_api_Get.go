// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package getinfo

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Get is the handler for GET /getinfo
// Send status information request
func (api GetinfoAPI) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var respBody types.GetinfoGetRespBody
	json.NewEncoder(w).Encode(&respBody)
}
