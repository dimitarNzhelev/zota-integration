package main

import (
	"fmt"
	"os"
	"zota_integration/Implementations"
	"zota_integration/structs"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	merchantId := os.Getenv("merchantId")
	merchantSecretKey := os.Getenv("merchantSecretKey")
	endpointID := os.Getenv("endpointID")
	url := os.Getenv("url")

	merchant := Implementations.NewMerchant(endpointID, url, merchantId, merchantSecretKey)

	res := merchant.Deposit(&structs.DepositPayload{
		MerchantOrderID:     "ABC123",
		MerchantOrderDesc:   "123",
		OrderAmount:         "123",
		OrderCurrency:       "USD",
		CustomerEmail:       "customer@email-address.com",
		CustomerFirstName:   "John",
		CustomerLastName:    "Doe",
		CustomerAddress:     "5/5 Moo 5 Thong Nai Pan Noi Beach, Baan Tai, Koh Phangan",
		CustomerCountryCode: "TH",
		CustomerCity:        "Surat Thani",
		CustomerZipCode:     "84280",
		CustomerPhone:       "+66-77999110",
		CustomerIP:          "103.106.8.104",
		RedirectURL:         "https://www.example-merchant.com/payment-return/",
		CallbackURL:         "https://www.example-merchant.com/payment-callback/",
		CheckoutURL:         "{\"UserId\": \"e139b447\"}",
		CustomParam:         "https://www.example-merchant.com/account/deposit/?uid=e139b447",
	})

	fmt.Println(res)

	res2 := merchant.Status(&structs.StatusPayload{
		OrderID:         "32529544",
		MerchantOrderID: "ABC123",
		MerchantID:      merchantId,
	})

	fmt.Println(res2)
}
