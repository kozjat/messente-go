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
		Text:           "Your SMS content here",
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("SMS sent! ID: %s\n", sentSms)
}
