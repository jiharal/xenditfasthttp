package xenditfasthttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// CoreXendit is ...
type CoreXendit struct {
	Client Client
}

const (
	pathCreateVA = "/callback_virtual_accounts"
)

// Call is ...
func (gw *CoreXendit) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}
	path = gw.Client.Host + path
	return gw.Client.Call(method, path, header, body, v)
}

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
	err = gw.Call(http.MethodPost, pathCreateVA, headers, buf, &res)
	if err != nil {
		return
	}
	return
}
