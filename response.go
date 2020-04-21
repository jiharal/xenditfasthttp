package xenditfasthttp

type (

	// CreateFixedPaymentCodeResponse is used to ....
	CreateFixedPaymentCodeResponse struct {
		ID               string `json:"id"`
		OwnerID          string `json:"owner_id"`
		ExternalID       string `json:"external_id"`
		RetailOutletName string `json:"retail_outlet_name"`
		Prefix           string `json:"prefix"`
		Name             string `json:"name"`
		PaymentCode      string `json:"payment_code"`
		Type             string `json:"type"`
		ExpectedAmount   int    `json:"expected_amount"`
		IsSingleUse      bool   `json:"is_single_use"`
		ExpirationDate   string `json:"expiration_date"`
	}

	// CreateVAResponse is used to...
	CreateVAResponse struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Status         string `json:"status"`
		Currency       string `json:"currency"`
		OwnerID        string `json:"owner_id"`
		ExternalID     string `json:"external_id"`
		BankCode       string `json:"bank_code"`
		MerchantCode   string `json:"merchant_code"`
		AccountNumber  string `json:"account_number"`
		ExpirationDate string `json:"expiration_date"`
		IsClosed       bool   `json:"is_closed"`
		IsSinglUse     bool   `json:"is_single_use"`
		Description    string `json:"description"`
	}
	// ResponseError is used to..
	ResponseError struct {
		ErrorCode string `json:"error_code"`
		Message   string `json:"message"`
	}
)
