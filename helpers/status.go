package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"zota_integration/structs"
)

func CheckStatus(status *structs.StatusPayload, minURL string) structs.Response {

	baseURL := minURL + "/query/order-status/"
	params := url.Values{}
	params.Add("merchantID", status.MerchantID)
	params.Add("merchantOrderID", status.MerchantOrderID)
	params.Add("orderID", status.OrderID)
	params.Add("timestamp", status.Timestamp)
	params.Add("signature", status.Signature)

	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var statusResponse structs.Response
	err = json.Unmarshal(body, &statusResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling response body: %v", err)
	}

	return statusResponse
}

func IsFinalStatus(status string) bool {
	fmt.Println(status)
	return status == "APPROVED" || status == "DECLINED" || status == "ERROR"
}
