package xenditfasthttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateFixedPaymentCode is used to create payment to outlet e.g: Alfamart, Indomaret ...
func (gw *CoreXendit) CreateFixedPaymentCode(req CreateFixedPaymentCodeRequest) (res CreateFixedPaymentCodeResponse, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	fmt.Println(buf.String())
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(http.MethodPost, patFixedPaymentCode, headers, buf, &res)
	if err != nil {
		return
	}
	return
}
