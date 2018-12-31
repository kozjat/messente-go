package sms

import (
	"errors"
	"fmt"
	"strings"

	"github.com/arvosaalits/messente-go/helpers"
)

// Arguments struct
type Arguments struct {
	SenderName     string
	ReceiverNumber string
}

// Send method sends the sms and returns boolean
func Send(config *Arguments) (string, error) {
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
