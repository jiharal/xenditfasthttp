package xenditfasthttp

type (
	// ResponseError is ..
	ResponseError struct {
		ErrorCode string `json:"error_code"`
		Message   string `json:"message"`
	}
	// CreateVAResponse is ...
	CreateVAResponse struct {
		IsClosed       bool   `json:"is_closed"`
		Status         string `json:"status"`
		Currency       string `json:"currency"`
		OwnerID        string `json:"owner_id"`
		ExternalID     string `json:"external_id"`
		BankCode       string `json:"bank_code"`
		MerchantCode   string `json:"merchant_code"`
		Name           string `json:"name"`
		AccountNumber  string `json:"account_number"`
		IsSinglUse     bool   `json:"is_single_use"`
		ExpirationDate string `json:"expiration_date"`
		ID             string `json:"id"`
	}
)
