package Interfaces

import "zota_integration/structs"

type Merchant interface {
	GenerateDepositSignature(merchantOrderID string, orderAmount string, customerEmail string) string
	GenerateStatusSignature(orderId string, merchantOrderId, timestamp string) string
	Deposit(deposit *structs.DepositPayload) structs.Response
	Status(status *structs.StatusPayload) structs.Response
	GetOrdersId() []string
}
