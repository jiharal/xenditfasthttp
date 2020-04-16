package xenditfasthttp

import (
	"fmt"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	auth := BasicAuth("", "")
	t.Log(auth)

	client := NewClient()
	client.Host = "https://api.xendit.co"
	client.SecretKey = ""

	core := CoreXendit{
		Client: client,
	}

	reqData := CreateVARequest{
		ExternalID: fmt.Sprintf("%v", time.Now().UnixNano()),
		BankCode:   VAPermata,
		Name:       "Jihar Al G",
		IsSinglUse: true,
	}

	resp, err := core.CreataVA(reqData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
