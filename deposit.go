package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zota_integration/structs"
)

type Response struct {
	Code    string
	Message string
	Data    map[string]interface{} `json:"data"`
}

var mockedDepositResult *Response

// this is for testing
func (mock *Response) SetMockResponse() {
	mockedDepositResult = mock
}

func (m *MerchantStruct) Deposit(d structs.DepositPayload) (res Response, err error) {

	err = d.Validate()
	if err != nil {
		return
	}

	err = m.validate()
	if err != nil {
		return
	}

	//this is for testing
	if mockedDepositResult != nil {
		res = *mockedDepositResult
		mockedDepositResult = nil
		return
	}

	d.Signature = m.genrateSignature(m.endpointID, d.MerchantOrderID, d.OrderAmount, d.CustomerEmail)

	deposit, err := json.Marshal(d)
	if err != nil {
		return
	}

	_, body, err := m.makeHttpReq(http.MethodPost, fmt.Sprintf("%v/api/v1/deposit/request/%v/", m.url, m.endpointID), deposit)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		err = fmt.Errorf("json Unmarshal err:%v", err)
		return
	}

	return
}


