// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package types

import (
	"gopkg.in/validator.v2"
)

type AddinvoicePostReqBody struct {
	Amt int `json:"amt" validate:"nonzero"`
}

func (s AddinvoicePostReqBody) Validate() error {

	return validator.Validate(s)
}
