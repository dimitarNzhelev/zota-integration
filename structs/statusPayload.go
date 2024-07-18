package structs

import (
	"fmt"
	"reflect"
)

type StatusPayload struct {
	MerchantID      string `json:"merchantID" validate:"required,max=32"`
	OrderID         string `json:"orderID" validate:"required,max=128"`
	MerchantOrderID string `json:"merchantOrderID" validate:"required,max=128"`
	Timestamp       string `json:"timestamp" validate:"required,max=15"`
	Signature       string `json:"signature" validate:"required,max=64"`
}

func (status *StatusPayload) Validate() error {
	required := []string{"MerchantOrderID", "OrderID"}
	for _, fieldName := range required {
		r := reflect.ValueOf(status)
		f := reflect.Indirect(r).FieldByName(fieldName)
		value := f.String()
		if value == "" {
			return fmt.Errorf("%v is required", fieldName)
		}
	}
	return nil
}
