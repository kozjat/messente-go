## API Wrapper for Messente (messente.com) APIs

### Request pricing list
```
resp, err := pricing.Listing("countryCode", "apiUsername", "apiPassword")
if err != nil {
    panic(err)
}

fmt.Println(resp) // Returns pricing.Response object
```
