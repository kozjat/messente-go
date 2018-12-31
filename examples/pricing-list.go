package examples

import (
	"fmt"

	"github.com/arvosaalits/messente-go/pricing"
)

// GetPricingList example
func GetPricingList() {
	provider, err := pricing.Listing("GB") // United Kingdom (GB) pricing list
	if err != nil {
		panic(err)
	}

	fmt.Println(provider) // pricing.Response object
}
