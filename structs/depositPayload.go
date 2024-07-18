package structs

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type DepositPayload struct {
	MerchantOrderID           string `json:"merchantOrderID"`
	MerchantOrderDesc         string `json:"merchantOrderDesc" `
	OrderAmount               string `json:"orderAmount"`
	OrderCurrency             string `json:"orderCurrency"`
	CustomerEmail             string `json:"customerEmail"`
	CustomerFirstName         string `json:"customerFirstName"`
	CustomerLastName          string `json:"customerLastName"`
	CustomerAddress           string `json:"customerAddress"`
	CustomerCountryCode       string `json:"customerCountryCode"`
	CustomerCity              string `json:"customerCity"`
	CustomerState             string `json:"customerState"`
	CustomerZipCode           string `json:"customerZipCode"`
	CustomerPhone             string `json:"customerPhone"`
	CustomerIP                string `json:"customerIP"`
	CustomerPersonalID        string `json:"customerPersonalID"`
	CustomerBankCode          string `json:"customerBankCode"`
	CustomerBankAccountNumber string `json:"customerBankAccountNumber"`
	RedirectURL               string `json:"redirectUrl"`
	CallbackURL               string `json:"callbackUrl"`
	CheckoutURL               string `json:"checkoutUrl"`
	CustomParam               string `json:"customParam"`
	Language                  string `json:"language"`
	Signature                 string `json:"signature"`
}

func (d *DepositPayload) Validate() error {
	validate := validator.New()

	err := validate.Struct(d)
	if err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	required := []string{"MerchantOrderID", "MerchantOrderDesc", "OrderAmount", "OrderCurrency", "MerchantOrderDesc", "CustomerEmail", "CustomerLastName", "CustomerAddress", "CustomerCountryCode", "CustomerCity", "CustomerZipCode", "CustomerPhone", "CustomerIP", "RedirectURL", "CheckoutURL"}
	for _, field := range required {
		if len(field) == 0 {
			return fmt.Errorf("validation error: %s is required", field)
		}
	}

	return nil
}
