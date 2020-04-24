package xenditfasthttp

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	pathDisbursements                = "/disbursements"
	pathAvailDisbBanks               = "/available_disbursements_banks"
	pathGetDisbursementByID          = "/disbursements/%s"
	pathGetDisbursementsByExternalID = "/disbursements?external_id=%s"
)

// GetDisbursementsByExternalID is used to get Disbursement by external ID
func (gw *CoreXendit) GetDisbursementsByExternalID(exID string) (res []DisbursementResponse, err error) {
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodGet, fmt.Sprintf(pathGetDisbursementsByExternalID, exID), headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// GetDisbursementByID is used to get disbusement by ID
func (gw *CoreXendit) GetDisbursementByID(disursementID string) (res DisbursementResponse, err error) {
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodGet, fmt.Sprintf(pathGetDisbursementByID, disursementID), headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// GetAvailableDisbursementBanks is used to get available disbusrsement banks
func (gw *CoreXendit) GetAvailableDisbursementBanks() (res []AvailableDisbursementBanksResponse, err error) {
	headers := map[string]string{
		"Authorization": BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":  "application/json",
	}
	err = gw.Call(fasthttp.MethodGet, pathAvailDisbBanks, headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// CreateDisbursement is used to create disbursement
func (gw *CoreXendit) CreateDisbursement(req DisbursementRequest) (res DisbursementResponse, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Authorization":     BasicAuth(gw.Client.SecretKey, ""),
		"Content-Type":      "application/json",
		"X-IDEMPOTENCY-KEY": req.IdempotencyKey,
	}
	err = gw.Call(fasthttp.MethodPost, pathDisbursements, headers, buf, &res)
	if err != nil {
		return
	}
	return
}
