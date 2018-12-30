## API Wrapper for Messente (messente.com) APIs

### Getting started
API access credentials should be created as environment variable
```
MESSENTE_API_USERNAME=your_api_key
MESSENTE_API_PASSWORD=your_api_secret
```

### [Pricing API] Request pricing list
```
resp, err := pricing.Listing("countryCode")
if err != nil {
    panic(err)
}

fmt.Println(resp) // pricing.Response object
```

### [Account Balance API] Get account balance
```
balance, err := balance.Get()
if err != nil {
    panic(err)
}

fmt.Println("Balance:", balance) // float64
```