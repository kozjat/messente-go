package examples

import (
	"fmt"

	"github.com/kozjat/messente-go/pricing"
)

// GetPricingList example
func GetPricingList() {
	provider, err := pricing.Listing("GB") // United Kingdom (GB) pricing list
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pricing list / Country name: %s, code: %s\n", provider.Name, provider.Country)
	for _, item := range provider.Networks {
		fmt.Printf("Provider name: %s, price (per message): %s\n", item.Provider, item.Price)
	}
}
