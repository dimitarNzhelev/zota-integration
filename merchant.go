package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

type MerchantStruct struct {
	endpointID        string
	url               string
	merchantId        string
	merchantSecretKey string
	HttpClient        httpClient
}

type httpClient interface {
	Do(req *http.Request) (ret *http.Response, err error)
}

func (m *MerchantStruct) validate() error {
	if m.endpointID == "" || m.url == "" || m.merchantId == "" || m.merchantSecretKey == "" {
		return fmt.Errorf("missing required fields")
	}
	return nil
}

func (m *MerchantStruct) genrateSignature(args ...string) (signature string) {

	str := ""

	for _, v := range args {
		str += v
	}

	h := sha256.New()
	h.Write([]byte(str + m.merchantSecretKey))

	signature = hex.EncodeToString(h.Sum(nil))
	return
}

func (m *MerchantStruct) initHttpClient() {
	if m.HttpClient == nil {
		m.HttpClient = &http.Client{}
	}
}

func (s *MerchantStruct) makeHttpReq(method string, url string, data []byte) (code int, body []byte, err error) {

	s.initHttpClient()

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.HttpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}

	code = resp.StatusCode
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
