package examples

import (
	"fmt"

	"github.com/kozjat/messente-go/balance"
)

// GetBalance example
func GetBalance() {
	balance, err := balance.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println("Account balance:", balance)
}
