package examples

import (
	"fmt"

	lookup "github.com/kozjat/messente-go/number-lookup"
)

// NumberLookup example
func NumberLookup() {
	lookup, err := lookup.Lookup([]string{
		"+372123456789", // Replace it with actual phone number
	})

	if err != nil {
		panic(err)
	}

	for _, item := range lookup.Result {
		fmt.Printf("Phone number: %s, status: %s\n", item.Number, item.Status)
		fmt.Printf(
			"Current network name: %s, country prefix: %s, country code: %s\n\n",
			item.CurrentNetwork.Name, item.CurrentNetwork.CountryPrefix, item.CurrentNetwork.CountryCode,
		)
	}
}
