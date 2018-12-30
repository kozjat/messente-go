package helpers

import "os"

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
