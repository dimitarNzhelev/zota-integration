package structs

import (
	"fmt"
	"reflect"
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
	required := []string{"MerchantOrderID", "MerchantOrderDesc", "OrderAmount", "OrderCurrency", "MerchantOrderDesc", "CustomerEmail", "CustomerLastName", "CustomerAddress", "CustomerCountryCode", "CustomerCity", "CustomerZipCode", "CustomerPhone", "CustomerIP", "RedirectURL", "CheckoutURL"}
	for _, fieldName := range required {
		r := reflect.ValueOf(d)
		f := reflect.Indirect(r).FieldByName(fieldName)
		value := f.String()
		if value == "" {
			return fmt.Errorf("%v is required", fieldName)
		}
	}
	return nil
}
