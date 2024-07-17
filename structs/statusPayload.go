package structs

type StatusPayload struct {
	MerchantID      string `json:"merchantID" validate:"required,max=32"`
	OrderID         string `json:"orderID" validate:"required,max=128"`
	MerchantOrderID string `json:"merchantOrderID" validate:"required,max=128"`
	Timestamp       string `json:"timestamp" validate:"required,max=15"`
	Signature       string `json:"signature" validate:"required,max=64"`
}
