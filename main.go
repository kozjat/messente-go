package main

import (
	"fmt"

	"github.com/arvosaalits/messente-api-client/pricing"
)

func main() {
	resp, err := pricing.Listing("GB", "apiUsername", "apiPassword")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
