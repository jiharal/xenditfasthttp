package xenditfasthttp

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	pathCreateVA                        = "/callback_virtual_accounts"
	pathGetVirtualAccountAvailableBanks = "/available_virtual_account_banks"
	pathGetVirtualAccountRequest        = "/callback_virtual_accounts/%s"
	pathUpdateVirtualAccount            = "/callback_virtual_accounts/%s"
)

// UpdateVirtualAccount is used to update virtual account
func (gw *CoreXendit) UpdateVirtualAccount(req UpdateVirtualAccountsRequest, id string) (res CreateVAResponse, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodPatch, fmt.Sprintf(pathUpdateVirtualAccount, id), headers, buf, &res)
	if err != nil {
		return
	}
	return
}

// GetVirtualAccountRequest is used to get virtual account request.
func (gw *CoreXendit) GetVirtualAccountRequest(id string) (res CreateVAResponse, err error) {
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodGet, fmt.Sprintf(pathGetVirtualAccountRequest, id), headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// GetVirtualAccountAvailableBanks is used to get virtual account available banks
func (gw *CoreXendit) GetVirtualAccountAvailableBanks() (res VirtualAccountBanksResponse, err error) {
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodGet, pathGetVirtualAccountAvailableBanks, headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// CreataVA is used to create virtual account
func (gw *CoreXendit) CreataVA(req CreateVARequest) (res CreateVAResponse, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodPost, pathCreateVA, headers, buf, &res)
	if err != nil {
		return
	}
	return
}
