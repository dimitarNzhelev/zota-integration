package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"zota_integration/structs"
)

func CreateDeposit(deposit *structs.DepositPayload, url string) structs.Response {

	depositBody, depositErr := json.Marshal(deposit)
	if depositErr != nil {
		log.Fatalf("Error marshalling deposit payload: %v", depositErr)
	}

	endpointID := os.Getenv("endpointID")
	fullURL := url + "/deposit/request/" + endpointID

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(depositBody))
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

	var depositResponse structs.Response
	err = json.Unmarshal(body, &depositResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling response body: %v", err)
	}

	return depositResponse
}
