package xenditfasthttp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	client := NewClient()
	client.Host = "https://api.xendit.co"
	client.SecretKey = "xnd_development_P4qDfOss0OCpl8RtKrROHjaQYNCk9dN5lSfk+R1l9Wbe+rSiCwZ3jw=="

	core := CoreXendit{
		Client: client,
	}

	auth := BasicAuth(client.SecretKey, "")
	t.Log(auth)

	reqDis := DisbursementRequest{
		ExternalID:        "JIHAR1122",
		BankCode:          "BCA",
		AccountHolderName: "Jihar Al Gifari",
		AccountNumber:     "1234567890",
		Description:       "Hello",
		Amount:            120000,
	}

	createDisembursement, err := core.CreateDisbursement(reqDis)
	require.NoError(t, err)
	t.Log(createDisembursement)

	availDisBanks, err := core.GetAvailableDisbursementBanks()
	require.NoError(t, err)
	t.Log(availDisBanks)

	a, err := core.GetDisbursementsByExternalID("JIHAR1")
	require.NoError(t, err)
	t.Log(a)

	reqData := CreateVARequest{
		ExternalID:  fmt.Sprintf("%v", time.Now().UnixNano()),
		BankCode:    VABRI,
		Name:        "Jihar",
		IsSingleUse: true,
	}

	resp, err := core.CreataVA(reqData)
	require.NoError(t, err)
	t.Log(resp)

	reqOutlet := CreateFixedPaymentCodeRequest{
		ExternalID:       fmt.Sprintf("KS%v", time.Now().UnixNano()),
		RetailOutletName: "INDOMARET",
		Name:             "Jihar",
		ExpectedAmount:   100000,
		IsSingleUse:      true,
	}
	outletResp, err := core.CreateFixedPaymentCode(reqOutlet)
	require.NoError(t, err)
	t.Log(outletResp)
}
