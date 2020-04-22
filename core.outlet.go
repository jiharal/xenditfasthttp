package xenditfasthttp

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	patFixedPaymentCode        = "/fixed_payment_code"
	pathUpdateFixedPaymentCode = "/fixed_payment_code/%s"
)

// UpdateFixedPaymentCode i sused to update payment
func (gw *CoreXendit) UpdateFixedPaymentCode(paymentCode string, req UpdateFixedPaymentCodeRequest) (res FixedPaymentCodeResponse, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodPatch, fmt.Sprintf(pathUpdateFixedPaymentCode, paymentCode), headers, buf, &res)
	if err != nil {
		return
	}
	return
}

// CreateFixedPaymentCode is used to create payment to outlet e.g: Alfamart, Indomaret ...
func (gw *CoreXendit) CreateFixedPaymentCode(req CreateFixedPaymentCodeRequest) (res FixedPaymentCodeResponse, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodPost, patFixedPaymentCode, headers, buf, &res)
	if err != nil {
		return
	}
	return
}
