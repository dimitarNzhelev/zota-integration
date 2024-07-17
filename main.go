package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
)


func generateSignature(endpointID, merchantOrderID, orderAmount, customerEmail, merchantSecretKey string) string {
    data := endpointID + merchantOrderID + orderAmount + customerEmail + merchantSecretKey

    hash := sha256.Sum256([]byte(data))

    signature := hex.EncodeToString(hash[:])

    return signature
}


func createHTTPReq(getPayload, getUrl func() string) {
	req, err := http.NewRequest("POST", getUrl(), bytes.NewBuffer([]byte(getPayload())))
    if err != nil {
        log.Fatalf("Error creating request: %v", err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error making request: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %v", err)
    }

    // Print the response status and body
    fmt.Printf("Response status: %s\n", resp.Status)
    fmt.Printf("Response body: %s\n", body)
}


func main() {    

    endpointID := "402334"
	merchantOrderID := "testNeAAAAwId"
	orderAmount := "99.99"
	customerEmail := "dimitar.n.zhelev@gmail.com"
	merchantSecretKey := "866adddb-7b91-4b1b-82a2-364479e17486"

	signature := generateSignature(endpointID, merchantOrderID, orderAmount, customerEmail, merchantSecretKey)

    url := "https://api.zotapay-stage.com/api/v1/deposit/request/402334/"

	payload := fmt.Sprintf(`{
        "merchantOrderID": "%s",
        "merchantOrderDesc": "WWWWWWWWWWWWWWWWWWWWW",
        "orderAmount": "%s",
        "orderCurrency": "USD",
        "customerEmail": "%s",
        "customerFirstName": "Ivna",
        "customerLastName": "Aswqq",
        "customerAddress": "The Swan, Jungle St. 108",
        "customerCountryCode": "US",
        "customerCity": "Los Angeles",
        "customerState": "CA",
        "customerZipCode": "90015",
        "customerPhone": "+1 420-100-1000",
        "customerBankCode": "BBL",
        "customerIP": "134.201.250.130",
        "redirectUrl": "https://www.example-merchant.com/payment-return/",
        "callbackUrl": "https://www.example-merchant.com/payment-callback/",
        "customParam": "{\"UserId\": \"c8266d59\"}",
        "checkoutUrl": "https://www.example-merchant.com/account/deposit/?uid=c8266d59",
        "signature": "%s"
    }`, merchantOrderID, orderAmount, customerEmail, signature)


    
}
