package balance

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kozjat/messente-go/helpers"
)

// Get returns the account balance
func Get() (float64, error) {
	credentials := helpers.APICredentials()

	url := fmt.Sprintf("%s/get_balance/?username=%s&password=%s", helpers.MessenteAPIUrl, credentials.Username, credentials.Password)
	body, err := helpers.ReadBody(url)
	if err != nil {
		return 0, err
	}

	prefix := "OK "
	var value float64
	if strings.HasPrefix(string(body), prefix) {
		stringValue := strings.TrimPrefix(string(body), prefix)
		value, err = strconv.ParseFloat(stringValue, 32)
		if err != nil {
			return 0, err
		}
	}

	return value, nil
}
