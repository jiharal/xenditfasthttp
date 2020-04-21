package xenditfasthttp

type (
	// CreateFixedPaymentCodeRequest is used to ..
	CreateFixedPaymentCodeRequest struct {
		ExternalID       string `json:"external_id"`
		RetailOutletName string `json:"retail_outlet_name"`
		Name             string `json:"name"`
		ExpectedAmount   int    `json:"expected_amount"`
		IsSingleUse      bool   `json:"is_single_use"`
	}
	// CreateVARequest is ...
	CreateVARequest struct {
		ExternalID           string `json:"external_id"`
		BankCode             string `json:"bank_code"`
		Name                 string `json:"name"`
		VirtualAccountNumber string `json:"virtual_account_number"`
		SuggestedAmount      int    `json:"suggested_amount"`
		IsClosed             bool   `json:"is_closed"`
		ExpectedAmount       int    `json:"expected_amount"`
		ExpirationDate       string `json:"expiration_date"`
		IsSingleUse          bool   `json:"is_single_use"`
		Description          string `json:"description"`
	}
)
