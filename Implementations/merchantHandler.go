package Implementations

import (
	"crypto/sha256"
	"encoding/hex"
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
	merchantSecretKey string
	statusChecker     Interfaces.StatusChecker
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
    if !ok {
        return response;
    }

    statusResponse := m.Status(&structs.StatusPayload{
        OrderID:         orderID,
        MerchantOrderID: deposit.MerchantOrderID,
        MerchantID:      m.merchantId,
    })
    status, ok := statusResponse.Data["status"].(string)
    if ok && m.statusChecker.IsFinalStatus(status) {
        return statusResponse
    }

    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()
    timeout := time.After(120 * time.Second)

    for {
        select {
        case <-ticker.C:
            statusResponse := m.Status(&structs.StatusPayload{
                OrderID:         orderID,
                MerchantOrderID: deposit.MerchantOrderID,
                MerchantID:      m.merchantId,
            })
            status, ok := statusResponse.Data["status"].(string)
            if ok && m.statusChecker.IsFinalStatus(status) {
                return statusResponse
            }
        case <-timeout:
            return structs.Response{Data: map[string]interface{}{"error": "Timeout waiting for final status"}}
        }
    }
}

func (m *MerchantStruct) Status(status *structs.StatusPayload) structs.Response {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signature := m.GenerateStatusSignature(status.OrderID, status.MerchantOrderID, timestamp)
	status.Signature = signature
	status.Timestamp = timestamp
	return helpers.CheckStatus(status, m.url)
}

func NewMerchant(endpointID, url, merchantId, merchantSecretKey string, statusChecker Interfaces.StatusChecker) Interfaces.Merchant {
	m := MerchantStruct{endpointID: endpointID, url: url, merchantId: merchantId, merchantSecretKey: merchantSecretKey, statusChecker: statusChecker}
	return &m
}
