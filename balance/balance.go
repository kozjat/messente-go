package balance

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/arvosaalits/messente-go/helpers"
)

// Get returns the account balance
func Get() (float64, error) {
	credentials := helpers.APICredentials()

	url := fmt.Sprintf("%s/get_balance/?username=%s&password=%s", helpers.MessenteAPIUrl, credentials.Username, credentials.Password)
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
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
