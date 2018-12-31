package main

import (
	"fmt"

	"github.com/kozjat/messente-go/balance"
	"github.com/kozjat/messente-go/pricing"
	"github.com/kozjat/messente-go/sms"
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

	sentSms, err := sms.Send(&sms.Options{
		ReceiverNumber: "+372123456789",
		SenderName:     "Sender",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(sentSms)
}
