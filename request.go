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

	// BatchDisbursementRequest is used to create dibursement at the same time.
	BatchDisbursementRequest struct {
		Reference     string              `json:"reference,required"`
		Disbursements []BatchDisbursement `json:"disbursements,required"`
	}

	// BatchDisbursement is used to create batch disb
	BatchDisbursement struct {
		Amount            int      `json:"amount,required"`
		BankCode          string   `json:"bank_code,required"`
		BankAccountName   string   `json:"bank_account_name,required"`
		BankAccountNumber string   `json:"bank_account_number,required"`
		Description       string   `json:"description,required"`
		ExternalID        string   `json:"external_id"`
		EmailTo           []string `json:"email_to,omitempty"`
		EmailCC           []string `json:"email_cc,omitempty"`
		EmailBCC          []string `json:"email_bcc,omitempty"`
	}

	// DisbursementRequest is a Disbursement
	DisbursementRequest struct {
		IdempotencyKey    string   `json:"idempotency_key,omitempty"`
		ExternalID        string   `json:"external_id,required"`
		BankCode          string   `json:"bank_code,required"`
		AccountHolderName string   `json:"account_holder_name,required"`
		AccountNumber     string   `json:"account_number,required"`
		Description       string   `json:"description,required"`
		Amount            int      `json:"amount,required"`
		EmailTo           []string `json:"email_to,omitempty"`
		EmailCC           []string `json:"email_cc,omitempty"`
		EmailBCC          []string `json:"email_bcc,omitempty"`
	}

	// CreateFixedPaymentCodeRequest is used to ..
	CreateFixedPaymentCodeRequest struct {
		ExternalID       string `json:"external_id,required"`
		RetailOutletName string `json:"retail_outlet_name,required"`
		Name             string `json:"name,required"`
		ExpectedAmount   int    `json:"expected_amount,required"`
		PaymentCode      string `json:"payment_code,omitempty"`
		ExpirationDate   string `json:"expiration_date,omitempty"`
		IsSingleUse      bool   `json:"is_single_use,omitempty"`
	}

	// SimulateRequest ...
	SimulateRequest struct {
		ExternalID       string `json:"external_id"`
		RetailOutletName string `json:"retail_outlet_name"`
		PaymentCode      string `json:"payment_code"`
		TransferAmount   int    `json:"transfer_amount"`
	}

	// CreateVARequest is ...
	CreateVARequest struct {
		ExternalID           string `json:"external_id,required"`
		BankCode             string `json:"bank_code,required"`
		Name                 string `json:"name,required"`
		VirtualAccountNumber string `json:"virtual_account_number,omitempty"`
		SuggestedAmount      int    `json:"suggested_amount,omitempty"`
		IsClosed             bool   `json:"is_closed,omitempty"`
		ExpectedAmount       int    `json:"expected_amount,omitempty"`
		ExpirationDate       string `json:"expiration_date,omitempty"`
		IsSingleUse          bool   `json:"is_single_use,omitempty"`
		Description          string `json:"description,omitempty"`
	}
)
