package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"zota_integration/structs"
)

type OrderStatusResult struct {
	Code    string                 `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

var mockedOrderStatusResult *OrderStatusResult

// this is for testing
func (mock *OrderStatusResult) SetMockResponse() {
	mockedOrderStatusResult = mock
}

func (m *MerchantStruct) Status(status structs.StatusPayload) (res OrderStatusResult, err error) {

	err = status.Validate()
	if err != nil {
		return
	}

	err = m.validate()
	if err != nil {
		return
	}

	status.MerchantID = m.merchantId

	//generate timestamp in unix format
	status.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)

	//only for testing
	if mockedOrderStatusResult != nil {
		res = *mockedOrderStatusResult
		mockedOrderStatusResult = nil
		return
	}

	status.Signature = m.genrateSignature(status.MerchantID, status.MerchantOrderID, status.OrderID, status.Timestamp)

	//is there a better way to do this?
	params := url.Values{}
	params.Add("merchantID", status.MerchantID)
	params.Add("merchantOrderID", status.MerchantOrderID)
	params.Add("orderID", status.OrderID)
	params.Add("timestamp", status.Timestamp)
	params.Add("signature", status.Signature)

	_, body, err := m.makeHttpReq(http.MethodGet, fmt.Sprintf("%v/api/v1/query/order-status/?%v", m.url, params.Encode()), []byte(""))
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
