// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package auth

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Post is the handler for POST /auth
// Send user authentication request
func (api AuthAPI) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody types.AuthPostReqBody

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var respBody types.AuthPostRespBody
	json.NewEncoder(w).Encode(&respBody)
}
