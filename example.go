package main

import (
	"fmt"
	"os"
	"zota_integration/merchant"
	"zota_integration/structs"
	"zota_integration/testutils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	m := merchant.NewMerchantStruct(
		os.Getenv("ENDPOINT_ID"),
		os.Getenv("URL"),
		os.Getenv("MERCHANT_ID"),
		os.Getenv("MERCHANT_SECRET_KEY"),
	)

	payload := testutils.NewRandomDepositPayload()

	res, err := m.Deposit(*payload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	statusPayload := structs.StatusPayload{
		MerchantOrderID: res.Data["merchantOrderID"].(string),
		OrderID:         res.Data["orderID"].(string),
	}

	statusRes, err := m.Status(statusPayload)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status Payload Result", statusRes)

}
