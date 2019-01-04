package sms

import "errors"

// ErrorCodes variable
var ErrorCodes = map[string]error{
	"ERROR 101":  errors.New("[SMS API] Access is restricted, wring credentials"),
	"ERROR 102":  errors.New("[SMS API] Parameters are wrong (or missing)"),
	"ERROR 103":  errors.New("[SMS API] Invalid IP address - not in the whitelist"),
	"ERROR 111":  errors.New("[SMS API] Sender parameter \"from\" is invalid"),
	"FAILED 209": errors.New("[SMS API] Server failure. Try again after a few seconds or try the api3.messente.com backup server"),
}
