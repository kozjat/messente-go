## API Wrapper for Messente (messente.com) APIs

### Getting started
API access credentials should be created as environment variable
```
MESSENTE_API_USERNAME=your_api_key
MESSENTE_API_PASSWORD=your_api_secret
```

### Request pricing list
```
resp, err := pricing.Listing("countryCode")
if err != nil {
    panic(err)
}

fmt.Println(resp) // Returns pricing.Response object
```
