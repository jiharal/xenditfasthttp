package xenditfasthttp

import (
	"bytes"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

const (
	pathCreateBathcDisbursement = "/batch_disbursements"
	pathGetAvailableBanks       = "/available_disbursements_banks"
)

// CreateBatchDisbursement is used to create batch disbursement
func (gw *CoreXendit) CreateBatchDisbursement(req BatchDisbursementRequest) (res CreateBatchDisbursementResponse, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodPost, pathCreateBathcDisbursement, headers, buf, &res)
	if err != nil {
		return
	}
	return
}

// GetAvailableBatchDisbursementBanks is used to get available banks.
func (gw *CoreXendit) GetAvailableBatchDisbursementBanks() (res []GetAvailableDisbursementBanksResponse, err error) {
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodGet, pathGetAvailableBanks, headers, nil, &res)
	if err != nil {
		return
	}
	return
}
