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

	// Text to send via SMS
	Text string
}

// Send method sends the request to  and returns the string that contains SMS ID
func Send(config *Options) (string, error) {
	credentials := helpers.APICredentials()

	url := fmt.Sprintf(
		"%s/send_sms/?username=%s&password=%s&from=%s&to=%s&text=%s",
		helpers.MessenteAPIUrl, credentials.Username, credentials.Password, config.SenderName, config.ReceiverNumber, config.Text,
	)

	body, err := helpers.Request("GET", url, []byte(""))
	if err != nil {
		return "", err
	}

	prefix := "OK "
	var value string
	if strings.HasPrefix(string(body), prefix) {
		value = strings.TrimPrefix(string(body), prefix)
	}

	err, ok := ErrorCodes[string(body)]
	if ok {
		return value, err
	}

	return value, errors.New("Unknown error" + string(body))
}
