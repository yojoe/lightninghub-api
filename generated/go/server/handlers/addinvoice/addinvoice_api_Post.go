// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package addinvoice

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Post is the handler for POST /addinvoice
// Request a BOLT-11 compatible Lightning invoice that can be paid by any other
// BOLT-11 compatible wallet. If the invoice is paid the respective amount
// will be credited to the authenticated user who created it.
func (api AddinvoiceAPI) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody types.AddinvoicePostReqBody

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var respBody types.AddinvoicePostRespBody
	json.NewEncoder(w).Encode(&respBody)
}
