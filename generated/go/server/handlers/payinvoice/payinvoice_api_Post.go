// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package payinvoice

import (
	"encoding/json"
	"github.com/yojoe/lightninghub-api/go/server/types"
	"net/http"
)

// Post is the handler for POST /payinvoice
// Request a BOLT-11 compatible Lightning invoice to be paid by the custodian.
// If the payment is successful, the respective amount + a fee is substracted
// from the authenticated user's balance.
func (api PayinvoiceAPI) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody types.PayinvoicePostReqBody

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var respBody types.PayinvoicePostRespBody
	json.NewEncoder(w).Encode(&respBody)
}
