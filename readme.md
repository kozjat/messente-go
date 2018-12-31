## WIP: API Wrapper for Messente (messente.com) APIs

### Getting started
API access credentials should be stored as an environment variables
```
MESSENTE_API_USERNAME=your_api_key
MESSENTE_API_PASSWORD=your_api_secret
```

### [Pricing API] Request pricing list
```
resp, err := pricing.Listing("countryCode") // pricing.Response, error
```

### [Account Balance API] Get account balance
```
balance, err := balance.Get() // float64, error
```

### [SMS API] Send sms
```
sentSms, err := sms.Send(&sms.Options{
    ReceiverNumber: "+372123456789",
    SenderName:     "SenderNameHere",
}) // string, error
```