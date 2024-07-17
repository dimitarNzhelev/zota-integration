package Implementations

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
	"zota_integration/Interfaces"
	"zota_integration/helpers"
	"zota_integration/structs"
)

type MerchantStruct struct {
	endpointID        string
	url               string
	merchantId        string
	ordersId          []string //not shure if this is the best way to store the ordersId or if it is necessary to store it
	merchantSecretKey string
}

func (m *MerchantStruct) GenerateDepositSignature(merchantOrderID, orderAmount, customerEmail string) string {
	data := m.endpointID + merchantOrderID + orderAmount + customerEmail + m.merchantSecretKey
	hash := sha256.Sum256([]byte(data))
	signature := hex.EncodeToString(hash[:])
	return signature
}

func (m *MerchantStruct) GenerateStatusSignature(orderId, merchantOrderId, timestamp string) string {
	data := m.merchantId + merchantOrderId + orderId + timestamp + m.merchantSecretKey
	hash := sha256.Sum256([]byte(data))
	signature := hex.EncodeToString(hash[:])
	return signature
}

func (m *MerchantStruct) Deposit(deposit *structs.DepositPayload) structs.Response {
	signature := m.GenerateDepositSignature(deposit.MerchantOrderID, deposit.OrderAmount, deposit.CustomerEmail)
	deposit.Signature = signature
	response := helpers.CreateDeposit(deposit, m.url)
	orderID, ok := response.Data["orderID"].(string)
	if ok {
		m.ordersId = append(m.ordersId, orderID)
	}
	return response
}

func (m *MerchantStruct) Status(status *structs.StatusPayload) structs.Response {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signature := m.GenerateStatusSignature(status.OrderID, status.MerchantOrderID, timestamp)
	fmt.Println(signature)
	status.Signature = signature
	status.Timestamp = timestamp
	return helpers.CheckStatus(status, m.url)
}

func (m *MerchantStruct) GetOrdersId() []string {
	return m.ordersId
}

func NewMerchant(endpointID, url, merchantId, merchantSecretKey string) Interfaces.Merchant {
	m := MerchantStruct{endpointID: endpointID, url: url, merchantId: merchantId, merchantSecretKey: merchantSecretKey, ordersId: []string{}}
	return &m
}
