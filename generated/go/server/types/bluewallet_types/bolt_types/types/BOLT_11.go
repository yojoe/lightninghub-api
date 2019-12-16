// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package types

import ()

// bech32 encoded bitcoin payment request according to the [BOLT-11](https://github.com/lightningnetwork/lightning-rfc/blob/master/11-payment-encoding.md)
// specification.
type BOLT_11 string

func (s BOLT_11) Validate() error {

	return nil
}
