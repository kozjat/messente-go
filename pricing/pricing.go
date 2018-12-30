package pricing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
func Listing(countryCode string, apiUsername string, apiPassword string) (*Response, error) {
	url := fmt.Sprintf("https://api2.messente.com/prices/?username=%s&password=%s&country=%s", apiUsername, apiPassword, countryCode)
	fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(response.Status, response.StatusCode, string(body))

	var responseObject Response
	json.Unmarshal(body, &responseObject)

	return &responseObject, nil
}
