package structs

type DepositPayload struct {
	MerchantOrderID           string `json:"merchantOrderID" validate:"required,max=128"`
	MerchantOrderDesc         string `json:"merchantOrderDesc" validate:"required,max=128"`
	OrderAmount               string `json:"orderAmount" validate:"required,max=24"`
	OrderCurrency             string `json:"orderCurrency" validate:"required,max=3"`
	CustomerEmail             string `json:"customerEmail" validate:"required,max=50"`
	CustomerFirstName         string `json:"customerFirstName" validate:"required,max=128"`
	CustomerLastName          string `json:"customerLastName" validate:"required,max=128"`
	CustomerAddress           string `json:"customerAddress" validate:"required,max=128"`
	CustomerCountryCode       string `json:"customerCountryCode" validate:"required,max=2"`
	CustomerCity              string `json:"customerCity" validate:"required,max=128"`
	CustomerState             string `json:"customerState" validate:"max=3"`
	CustomerZipCode           string `json:"customerZipCode" validate:"required,max=15"`
	CustomerPhone             string `json:"customerPhone" validate:"required,max=15"`
	CustomerIP                string `json:"customerIP" validate:"required,max=64"`
	CustomerPersonalID        string `json:"customerPersonalID" validate:"max=20"`
	CustomerBankCode          string `json:"customerBankCode" validate:"max=16"`
	CustomerBankAccountNumber string `json:"customerBankAccountNumber" validate:"max=64"`
	RedirectURL               string `json:"redirectUrl" validate:"required,max=255"`
	CallbackURL               string `json:"callbackUrl" validate:"max=255"`
	CheckoutURL               string `json:"checkoutUrl" validate:"required,max=256"`
	CustomParam               string `json:"customParam" validate:"max=128"`
	Language                  string `json:"language" validate:"max=2"`
	Signature                 string `json:"signature" validate:"required,max=64"`
}
