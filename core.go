package xenditfasthttp

import (
	"fmt"
	"io"
	"strings"
)

// CoreXendit is ...
type CoreXendit struct {
	Client Client
}

// Call is ...
func (gw *CoreXendit) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}
	path = gw.Client.Host + path
	return gw.Client.Call(method, path, header, body, v)
}
