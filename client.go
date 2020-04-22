package xenditfasthttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

// Client is ...
type Client struct {
	Host      string
	SecretKey string
	LogLevel  int
	Logger    *log.Logger
}

var (
	defHTTPTimeout = 15 * time.Second
)

// NewClient is ...
func NewClient() Client {
	return Client{
		LogLevel: 2,
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
	}
}

// NewRequest test is ...
func (c *Client) NewRequest(method, fullPath string, headers map[string]string, body io.Reader) (*fasthttp.Request, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fullPath)
	req.Header.SetMethod(method)
	if method == fasthttp.MethodPost {
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)
		req.SetBody(buf.Bytes())
	}
	if headers != nil {
		for k, vv := range headers {
			req.Header.Set(k, vv)
		}
	}
	return req, nil
}

// ExecuteRequest is ...
func (c *Client) ExecuteRequest(req *fasthttp.Request, v interface{}) error {
	logLevel := c.LogLevel
	logger := c.Logger
	if logLevel > 1 {
		logger.Printf("Request %s:%s%s", string(req.Header.Method()), string(req.Host()), string(req.URI().Path()))
	}
	start := time.Now()
	resp := fasthttp.AcquireResponse()
	httpClient := &fasthttp.Client{}
	err := httpClient.Do(req, resp)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot send request: ", err)
		}
		return err
	}
	if logLevel > 2 {
		logger.Println("Completed in ", time.Since(start))
	}
	if v != nil && resp.StatusCode() == 200 {
		if err = json.Unmarshal(resp.Body(), v); err != nil {
			return err
		}
		return nil
	}
	var respErr ResponseError
	if err = json.Unmarshal(resp.Body(), &respErr); err != nil {
		return err
	}
	return fmt.Errorf("%s-%s", respErr.ErrorCode, respErr.Message)
}

// Call is ...
func (c *Client) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, header, body)
	if err != nil {
		return err
	}
	return c.ExecuteRequest(req, v)
}
