package logger

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

// StartRecord for initialize context first time
func StartRecord(req *http.Request, start time.Time) *http.Request {
	var body string

	if req.Method != http.MethodGet && !strings.Contains(req.Header.Get("Content-Type"), "form-data") {
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			body = ""
		} else {
			body = string(reqBody)
			req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		}
	}

	ctx := req.Context()

	v := new(Data)
	v.RequestID = uuid.New().String()
	v.Service = filepath.Base(os.Args[0])
	v.Host = req.Host
	v.Endpoint = req.URL.Path
	v.TimeStart = start
	v.Device = services
	v.UserCode = services
	v.RequestHeader = DumpRequest(req)
	v.RequestBody = body

	ctx = context.WithValue(ctx, logKey, v)

	return req.WithContext(ctx)
}

// DumpRequest is for get all data request header
func DumpRequest(req *http.Request) string {
	header, err := httputil.DumpRequest(req, false)
	if err != nil {
		return "cannot dump request"
	}

	trim := bytes.ReplaceAll(header, []byte("\r\n"), []byte("   "))
	return string(trim)
}
