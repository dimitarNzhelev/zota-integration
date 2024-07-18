package merchant

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

//String Method
func (r OrderStatusResult) String() string {
	return fmt.Sprintf("Code: %v, Message: %v, Data: %v", r.Code, r.Message, r.Data)
}

var mockedOrderStatusResult *OrderStatusResult

// this is for testing
func (mock *OrderStatusResult) SetMockResponse() {
	mockedOrderStatusResult = mock
}

// Status method (core method)
func (m *MerchantStruct) Status(status structs.StatusPayload) (res OrderStatusResult, err error) {

	err = m.validate()
	if err != nil {
		return
	}

	status.MerchantID = m.MerchantID
	//generate timestamp in unix format
	status.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	status.Signature = m.genrateSignature(status.MerchantID, status.MerchantOrderID, status.OrderID, status.Timestamp)

	err = status.Validate()
	if err != nil {
		return
	}

	//only for testing
	if mockedOrderStatusResult != nil {
		res = *mockedOrderStatusResult
		mockedOrderStatusResult = nil
		return
	}


	// create the query params
	params := url.Values{}
	params.Add("merchantID", status.MerchantID)
	params.Add("merchantOrderID", status.MerchantOrderID)
	params.Add("orderID", status.OrderID)
	params.Add("timestamp", status.Timestamp)
	params.Add("signature", status.Signature)

	_, body, err := m.PerformRequest(http.MethodGet, fmt.Sprintf("%v/api/v1/query/order-status/?%v", m.Url, params.Encode()), []byte(""))
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
