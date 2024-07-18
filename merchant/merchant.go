package merchant

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

type MerchantStruct struct {
	EndpointID        string
	Url               string
	MerchantID        string
	MerchantSecretKey string
	HttpClient        httpClient
	ReqMaker          HttpReqMaker
}

type RealReqMaker struct{}

type HttpReqMaker interface {
	makeHttpReq(client httpClient, method, url string, body []byte) (int, []byte, error)
}


type httpClient interface {
	Do(req *http.Request) (ret *http.Response, err error)
}

func (m *MerchantStruct) validate() error {
	if m.EndpointID == "" || m.Url == "" || m.MerchantID == "" || m.MerchantSecretKey == "" {
		return fmt.Errorf("validation error")
	}
	return nil
}

func (m *MerchantStruct) genrateSignature(args ...string) (signature string) {

	str := ""

	for _, v := range args {
		str += v
	}

	h := sha256.New()
	h.Write([]byte(str + m.MerchantSecretKey))

	signature = hex.EncodeToString(h.Sum(nil))
	return
}

// initHttpClient initializes the http client
func (m *MerchantStruct) initHttpClient() {
	if m.HttpClient == nil {
		m.HttpClient = &http.Client{}
	}
}

// makeHttpReq makes an http request
func (r RealReqMaker) makeHttpReq(client httpClient, method, url string, body []byte) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, responseBody, nil
}

// PerformRequest makes a request to the merchant server
func (m *MerchantStruct) PerformRequest(method string, url string, data []byte) (int, []byte, error) {
	m.initHttpClient()

	return m.ReqMaker.makeHttpReq(m.HttpClient, method, url, data)
}


// constructor
func NewMerchantStruct(endpointID, url, merchantID, merchantSecretKey string) *MerchantStruct {
	return &MerchantStruct{
		EndpointID:        endpointID,
		Url:               url,
		MerchantID:        merchantID,
		MerchantSecretKey: merchantSecretKey,
		ReqMaker:          RealReqMaker{},
	}
}
