package xenditfasthttp

import (
	"bytes"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// CreataVA is ...
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
