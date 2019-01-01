package lookup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kozjat/messente-go/helpers"
)

// Network struct
type Network struct {
	CounryName    string `json:"countryName"`
	Mccmnc        string `json:"mccmnc"`
	CountryPrefix string `json:"countryPrefix"`
	CountryCode   string `json:"countryCode"`
	Name          string `json:"networkName"`
}

// Result struct
type Result struct {
	Number          string  `json:"number"`
	CurrentNetwork  Network `json:"currentNetwork"`
	OriginalNetwork Network `json:"originalNetwork"`
	Ported          bool    `json:"ported"`
	PortedNetwork   Network `json:"portedNetwork"`
	Roaming         bool    `json:"roaming"`
	RoamingNetwork  Network `json:"roamingNetwork"`
	Status          string  `json:"status"`
}

// Response struct
type Response struct {
	Result []Result `json:"result"`
	Error  string   `json:"error"`
}

// RequestBody struct
type RequestBody struct {
	Numbers []string `json:"numbers"`
}

// Lookup method
func Lookup(numbers []string) (*Response, error) {
	credentials := helpers.APICredentials()

	uri := fmt.Sprintf("%s/hlr/sync", helpers.BaseAPIUrl)
	data, _ := json.Marshal(RequestBody{
		Numbers: numbers,
	})

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	req.SetBasicAuth(credentials.Username, credentials.Password)
	req.Header.Set("Content-Type", "application/json")

	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseObject Response
	json.Unmarshal(body, &responseObject)

	return &responseObject, nil
}
