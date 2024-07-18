package merchant

import (
	"fmt"
	"testing"
	"zota_integration/structs"
)

type MockReqMakerStatus struct{}

func TestStatus_Success(t *testing.T) {
    // Mock response
    mockResponse := &OrderStatusResult{
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

    statusPayload := structs.StatusPayload{
        MerchantID:      "merchant123",
        OrderID:         "order123",
        MerchantOrderID: "merchantOrder123",
        Timestamp:       "1609459200",
        Signature:       "signatureExample",
    }

    res, err := merchant.Status(statusPayload)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

	fmt.Println(res)

	if res.Code != "200" {
        t.Fatalf("expected success response, got %v", res)
    }
}

func TestStatus_ValidationError(t *testing.T) {
    // This test should simulate a validation error
    merchant := MerchantStruct{
        Url:               "https://example.com",
        EndpointID:        "testEndpoint",
        MerchantID:        "merchant123",
        MerchantSecretKey: "secret123",
    }

    statusPayload := structs.StatusPayload{}

    _, err := merchant.Status(statusPayload)
    if err == nil {
        t.Fatalf("expected validation error, got nil")
    }
}