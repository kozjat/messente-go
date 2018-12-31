package examples

import (
	"fmt"

	"github.com/kozjat/messente-go/sms"
)

// SendSMS example
func SendSMS() {
	sentSms, err := sms.Send(&sms.Options{
		ReceiverNumber: "+372123456789",
		SenderName:     "SenderNameHere",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(sentSms)
}
