package xenditfasthttp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateDisbursement(t *testing.T) {
	client := NewClient()
	client.Host = "https://api.xendit.co"
	client.SecretKey = ""

	core := CoreXendit{
		Client: client,
	}

	getAvailBatch, err := core.GetAvailableBatchDisbursementBanks()
	require.NoError(t, err)
	t.Log(getAvailBatch)

	batchDidbReq := BatchDisbursementRequest{
		Reference: fmt.Sprintf("withdraw_%v", time.Now().Unix()),
		Disbursements: []BatchDisbursement{
			{
				ExternalID:        fmt.Sprintf("kesan_%v", time.Now().Unix()),
				BankCode:          "BCA",
				Amount:            90000,
				BankAccountName:   "Joe",
				BankAccountNumber: "1234567890",
				Description:       "Batch disbursement",
				EmailTo:           []string{"jihar@tlab.co.id"},
			},
		},
	}
	resp, err := core.CreateBatchDisbursement(batchDidbReq)
	require.NoError(t, err)
	t.Log(resp)
	reqData := DisbursementRequest{
		IdempotencyKey:    fmt.Sprintf("key_%v", time.Now().Unix()),
		ExternalID:        fmt.Sprintf("kesan_%v", time.Now().Unix()),
		BankCode:          "BCA",
		Amount:            15000,
		AccountHolderName: "Joe",
		AccountNumber:     "1234567890",
		Description:       "Disbursement from postman",
	}

	res, err := core.CreateDisbursement(reqData)
	require.NoError(t, err)
	t.Log(res)

	getAvailableDisbursementBanks, err := core.GetAvailableDisbursementBanks()
	require.NoError(t, err)
	t.Log(getAvailableDisbursementBanks)

	getDisbursementByID, err := core.GetDisbursementByID("5ea2509fac88430016567a59")
	require.NoError(t, err)
	t.Log(getDisbursementByID)

	getDisbursementsByExternalID, err := core.GetDisbursementsByExternalID("kesan_1587695773")
	require.NoError(t, err)
	t.Log(getDisbursementsByExternalID)
}
