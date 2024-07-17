package testutils

import (
	"math/rand"
	"strconv"
	"zota_integration/structs"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func NewRandomDepositPayload() *structs.DepositPayload {
	return &structs.DepositPayload{
		MerchantOrderID:     RandomString(10),
		MerchantOrderDesc:   RandomString(10),
		OrderAmount:         strconv.Itoa(RandomInt(1, 100)),
		OrderCurrency:       "USD",
		CustomerEmail:       RandomString(10) + "@example.com",
		CustomerFirstName:   RandomString(5),
		CustomerLastName:    RandomString(5),
		CustomerAddress:     "5/5 Moo 5 Thong Nai Pan Noi Beach, Baan Tai, Koh Phangan",
		CustomerCountryCode: "TH",
		CustomerCity:        "Surat Thani",
		CustomerZipCode:     "84280",
		CustomerPhone:       "+66-77999110",
		CustomerIP:          "103.106.8.104",
		RedirectURL:         "https://www.example-merchant.com/payment-return/",
		CustomParam:         "{\"UserId\": \"e139b447\"}",
		CheckoutURL:         "https://www.example-merchant.com/account/deposit/?uid=e139b447",
	}
}
