package examples

import (
	"fmt"

	"github.com/arvosaalits/messente-go/balance"
)

// GetBalance example
func GetBalance() {
	balance, err := balance.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println("Balance:", balance)
}
