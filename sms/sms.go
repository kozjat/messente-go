package sms

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kozjat/messente-go/helpers"
)

// Options is an options struct for the Send method
type Options struct {
	// Name that sms receiver see as a sender name
	SenderName string

	// Receiver number that starts with area code
	ReceiverNumber string
}

// Send method sends the sms and returns the string
func Send(config *Options) (string, error) {
	credentials := helpers.APICredentials()

	url := fmt.Sprintf(
		"%s/send_sms/?username=%s&password=%s&from=%s&to=%s",
		helpers.MessenteAPIUrl, credentials.Username, credentials.Password, config.SenderName, config.ReceiverNumber,
	)

	body, err := helpers.ReadBody(url)
	if err != nil {
		return "", err
	}

	prefix := "OK "
	var value string
	if strings.HasPrefix(body, prefix) {
		value = strings.TrimPrefix(body, prefix)
	}

	return value, errors.New(string(body))
}
