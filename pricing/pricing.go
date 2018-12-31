package pricing

import (
	"encoding/json"
	"fmt"

	"github.com/kozjat/messente-go/helpers"
)

// Network struct
type Network struct {
	Mccmnc   string `json:"mccmnc"`
	Provider string `json:"name"`
	Price    string `json:"price"`
}

// Response struct
type Response struct {
	Country  string    `json:"country"`
	Name     string    `json:"name"`
	Prefix   string    `json:"prefix"`
	Networks []Network `json:"networks"`
}

// Listing returns the list of the prices
func Listing(countryCode string) (*Response, error) {
	credentials := helpers.APICredentials()

	url := fmt.Sprintf("%s/prices/?username=%s&password=%s&country=%s", helpers.MessenteAPIUrl, credentials.Username, credentials.Password, countryCode)
	body, err := helpers.ReadBody(url)
	if err != nil {
		return nil, err
	}

	var responseObject Response
	json.Unmarshal([]byte(body), &responseObject)

	return &responseObject, nil
}
