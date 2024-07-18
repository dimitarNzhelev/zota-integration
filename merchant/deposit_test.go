package merchant

import (
	"testing"
	"zota_integration/structs"
	"zota_integration/testutils"
)

type MockReqMaker struct{}


func TestDeposit_Success(t *testing.T) {
	mockResponse := &Response{
		Code:    "200",
		Message: "",
		Data:    map[string]interface{}{},
	}
	mockResponse.SetMockResponse()

	merchant := MerchantStruct{
		Url:               "https://example.com",
		EndpointID:        "testEndpoint",
		MerchantID:        "merchant123",
		MerchantSecretKey: "secret123",
	}

	payload := testutils.NewRandomDepositPayload()

	res, err := merchant.Deposit(*payload)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res.Code != "200" {
		t.Fatalf("expected success response, got %v", res)
	}
}

func TestDeposit_ValidationError(t *testing.T) {
	merchant := MerchantStruct{
		Url:        "https://example.com",
		EndpointID: "testEndpoint",
	}

	// Payload with missing required fields to trigger validation error
	payload := structs.DepositPayload{
		MerchantOrderID:     "",
		MerchantOrderDesc:   "",
		OrderAmount:         "100.00",
		OrderCurrency:       "USD",
		CustomerEmail:       "invalid_email",
		CustomerFirstName:   "John",
		CustomerLastName:    "Doe",
		CustomerAddress:     "123 Street Name",
		CustomerCountryCode: "US",
		CustomerCity:        "CityName",
		CustomerZipCode:     "12345",
		CustomerPhone:       "1234567890",
		CustomerIP:          "192.168.1.1",
		RedirectURL:         "https://example.com/redirect",
		CallbackURL:         "https://example.com/callback",
		CheckoutURL:         "https://example.com/checkout",
		Signature:           "signature123",
	}

	_, err := merchant.Deposit(payload)
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}

	expectedErr := "validation error"
	if err.Error()[:len(expectedErr)] != expectedErr {
		t.Fatalf("expected validation error, got %v", err)
	}
}
