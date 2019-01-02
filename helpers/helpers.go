package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

// BaseAPIUrl is the base ur of the Messente's main API
const BaseAPIUrl = "https://api.messente.com"

// MessenteAPIUrl is the base url of the Messente's V2 API
const MessenteAPIUrl = "https://api2.messente.com"

// Credentials struct stores the username and password
type Credentials struct {
	Username string
	Password string
}

// APICredentials method returns the credentials object
func APICredentials() Credentials {
	return Credentials{
		Username: os.Getenv("MESSENTE_API_USERNAME"),
		Password: os.Getenv("MESSENTE_API_PASSWORD"),
	}
}

// Request method
func Request(method string, url string, data []byte) ([]byte, error) {
	payload := bytes.NewBuffer(data)
	req, _ := http.NewRequest(method, url, payload)

	if method == "POST" {
		credentials := APICredentials()
		req.SetBasicAuth(credentials.Username, credentials.Password)
		req.Header.Set("Content-Type", "application/json")
	}

	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
