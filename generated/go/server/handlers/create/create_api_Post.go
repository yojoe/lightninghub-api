// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package create

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Post is the handler for POST /create
// Send account creation request
func (api CreateAPI) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody types.CreatePostReqBody

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var respBody types.CreatePostRespBody
	json.NewEncoder(w).Encode(&respBody)
}
