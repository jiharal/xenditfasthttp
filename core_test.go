package xenditfasthttp

import (
	"fmt"
	"testing"
	"time"
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

	reqData := CreateVARequest{
		ExternalID:  fmt.Sprintf("%v", time.Now().UnixNano()),
		BankCode:    VABRI,
		Name:        "Jihar Al G",
		IsSingleUse: true,
	}

	resp, err := core.CreataVA(reqData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

	t.Log(fmt.Sprintf("KS%v", time.Now().UnixNano()))

	reqOutlet := CreateFixedPaymentCodeRequest{
		ExternalID:       fmt.Sprintf("KS%v", time.Now().UnixNano()),
		RetailOutletName: "INDOMARET",
		Name:             "Jihar",
		ExpectedAmount:   100000,
		IsSingleUse:      true,
	}
	outletResp, err := core.CreateFixedPaymentCode(reqOutlet)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(outletResp)
}
