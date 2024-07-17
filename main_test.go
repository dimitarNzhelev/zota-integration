package main

import (
	"fmt"
	"os"
	"testing"
	"zota_integration/Implementations"
	"zota_integration/helpers"
	"zota_integration/testutils"

	"github.com/joho/godotenv"
)

func TestCreationOfDeposit(t *testing.T) {
	godotenv.Load(".env")

	merchantId := os.Getenv("merchantId")
	merchantSecretKey := os.Getenv("merchantSecretKey")
	endpointID := os.Getenv("endpointID")
	url := os.Getenv("url")

	statusChecker := helpers.TesterStatusChecker{}

	merchant := Implementations.NewMerchant(endpointID, url, merchantId, merchantSecretKey, statusChecker)

	deposit := testutils.NewRandomDepositPayload()

	response := merchant.Deposit(deposit)

	fmt.Println(response)

	if response.Code != "200" && (response.Data["status"] != "CREATED" && response.Data["status"] != "APPROVED") {
		t.Errorf("Expected response status to be CREATED or APPROVED , got %s", response.Data["status"])
	}
}
