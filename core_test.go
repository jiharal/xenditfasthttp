package xenditfasthttp

import (
	"fmt"
	"testing"
	"time"
)

func initial() *CoreXendit {
	client := NewClient()
	client.Host = "https://api.xendit.co"
	client.SecretKey = ""

	return &CoreXendit{
		Client: client,
	}
}
func TestOutlet(t *testing.T) {
	init := initial()
	req := CreateFixedPaymentCodeRequest{
		ExternalID:       fmt.Sprintf("%v", time.Now().Unix()),
		RetailOutletName: OutletIndomaret,
		Name:             "Joe Doe",
		ExpectedAmount:   10000,
		ExpirationDate:   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		IsSingleUse:      true,
	}
	reqData, err := init.CreateFixedPaymentCode(req)
	if err != nil {
		t.Fatal(err)
	}

	reqSimulate := SimulateRequest{
		ExternalID:       reqData.ExternalID,
		RetailOutletName: reqData.RetailOutletName,
		PaymentCode:      reqData.PaymentCode,
		TransferAmount:   reqData.ExpectedAmount,
	}

	simResp, err := init.SimulateFixedPayment(reqSimulate)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(simResp)
}

func TestVA(t *testing.T) {
	init := initial()

	reqData := CreateVARequest{
		ExternalID:      fmt.Sprintf("%v", time.Now().UnixNano()),
		BankCode:        VABRI,
		Name:            "Jihar",
		IsSingleUse:     true,
		SuggestedAmount: 100000,
		ExpirationDate:  time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}
	resp, err := init.CreataVA(reqData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
