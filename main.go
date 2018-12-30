package main

import (
	"fmt"

	"github.com/arvosaalits/messente-go/balance"
	"github.com/arvosaalits/messente-go/pricing"
)

func main() {
	resp, err := pricing.Listing("GB")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	balance, err := balance.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println("Balance:", balance)
}
