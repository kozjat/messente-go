package main

import (
	"fmt"

	"github.com/arvosaalits/messente-go/balance"
	"github.com/arvosaalits/messente-go/pricing"
	"github.com/arvosaalits/messente-go/sms"
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

	sentSms, err := sms.Send(&sms.Arguments{
		ReceiverNumber: "+372123456789",
		SenderName:     "Sender",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(sentSms)
}
