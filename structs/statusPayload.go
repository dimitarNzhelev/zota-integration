package structs

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type StatusPayload struct {
	MerchantID      string `json:"merchantID" validate:"required,max=32"`
	OrderID         string `json:"orderID" validate:"required,max=128"`
	MerchantOrderID string `json:"merchantOrderID" validate:"required,max=128"`
	Timestamp       string `json:"timestamp" validate:"required,max=15"`
	Signature       string `json:"signature" validate:"required,max=64"`
}

func (status *StatusPayload) Validate() error {
	validate := validator.New()

	err := validate.Struct(status)
	if err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	required := []string{"MerchantOrderID", "OrderID"}
	for _, field := range required {
		if len(field) == 0 {
			return fmt.Errorf("validation error: %s is required", field)
		}
	}

	return nil
}
