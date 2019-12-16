// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package types

type EnumPaymentType string

const (
	EnumPaymentTypebitcoind_tx  EnumPaymentType = "bitcoind_tx"
	EnumPaymentTypepaid_invoice EnumPaymentType = "paid_invoice"
	EnumPaymentTypeuser_invoice EnumPaymentType = "user_invoice"
)
