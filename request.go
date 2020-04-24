package xenditfasthttp

type (
	// UpdateVirtualAccountsRequest is used to update virtual account
	UpdateVirtualAccountsRequest struct {
		SuggestedAmount int    `json:"suggested_amount,omitempty"`
		ExpectedAmount  int    `json:"expected_amount,omitempty"`
		ExpirationDate  string `json:"expiration_date,omitempty"`
		IsSingleUse     bool   `json:"is_single_use,omitempty"`
		Description     string `json:"description,omitempty"`
	}

	// UpdateFixedPaymentCodeRequest is used to update ...
	UpdateFixedPaymentCodeRequest struct {
		Name           string `json:"name,omitempty"`
		ExpectedAmount int    `json:"expected_amount,omitempty"`
		ExpirationDate string `json:"expiration_date,omitempty"`
	}

	// DisbursementRequest is a Disbursement
	DisbursementRequest struct {
		IdempotencyKey    string   `json:"idempotency_key"`
		ExternalID        string   `json:"external_id"`
		BankCode          string   `json:"bank_code"`
		AccountHolderName string   `json:"account_holder_name"`
		AccountNumber     string   `json:"account_number"`
		Description       string   `json:"description"`
		Amount            int      `json:"amount"`
		EmailTo           []string `json:"email_to,omitempty"`
		EmailCC           []string `json:"email_cc,omitempty"`
		EmailBCC          []string `json:"email_bcc,omitempty"`
	}

	// CreateFixedPaymentCodeRequest is used to ..
	CreateFixedPaymentCodeRequest struct {
		ExternalID       string `json:"external_id"`
		RetailOutletName string `json:"retail_outlet_name"`
		Name             string `json:"name"`
		ExpectedAmount   int    `json:"expected_amount"`
		PaymentCode      string `json:"payment_code,omitempty"`
		ExpirationDate   string `json:"expiration_date,omitempty"`
		IsSingleUse      bool   `json:"is_single_use,omitempty"`
	}
	// CreateVARequest is ...
	CreateVARequest struct {
		ExternalID           string `json:"external_id"`
		BankCode             string `json:"bank_code"`
		Name                 string `json:"name"`
		VirtualAccountNumber string `json:"virtual_account_number,omitempty"`
		SuggestedAmount      int    `json:"suggested_amount,omitempty"`
		IsClosed             bool   `json:"is_closed,omitempty"`
		ExpectedAmount       int    `json:"expected_amount,omitempty"`
		ExpirationDate       string `json:"expiration_date,omitempty"`
		IsSingleUse          bool   `json:"is_single_use,omitempty"`
		Description          string `json:"description,omitempty"`
	}
)
