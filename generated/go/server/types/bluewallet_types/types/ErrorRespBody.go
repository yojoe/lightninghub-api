// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package types

import (
	"gopkg.in/validator.v2"
)

type ErrorRespBody struct {
	Code    int    `json:"code" validate:"min=0,nonzero"`
	Error   bool   `json:"error"`
	Message string `json:"message" validate:"nonzero"`
}

func (s ErrorRespBody) Validate() error {

	return validator.Validate(s)
}