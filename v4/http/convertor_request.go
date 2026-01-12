package http

import (
	"fmt"
	"io"
	"net/http"

	types "github.com/spinframework/spin-go-sdk/v4/wit_component/wasi_http_0_3_0_rc_2025_09_16_types"
)

// convert the IncomingRequest to http.Request
func NewHttpRequest(req *types.Request) (*http.Request, error) {
	if req == nil {
		return nil, fmt.Errorf("Request can't be nil")
	}

	// convert the http method to string
	method, err := methodToString(req.GetMethod())
	if err != nil {
		return nil, err
	}

	// convert the path with query to a url
	var url string
	if pathWithQuery := req.GetPathWithQuery(); pathWithQuery.IsNone() {
		url = ""
	} else {
		url = pathWithQuery.Some()
	}

	// convert the body to a reader
	var body io.Reader
	// TODO

	// create a new request
	convertedRequest, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// update additional fields
	toHttpHeader(*req.GetHeaders(), &convertedRequest.Header)

	return convertedRequest, nil
}

func methodToString(m types.Method) (string, error) {
	switch m.Tag() {
	case types.MethodConnect:
		return "CONNECT", nil
	case types.MethodDelete:
		return "DELETE", nil
	case types.MethodGet:
		return "GET", nil
	case types.MethodHead:
		return "HEAD", nil
	case types.MethodOptions:
		return "OPTIONS", nil
	case types.MethodPatch:
		return "PATCH", nil
	case types.MethodPost:
		return "POST", nil
	case types.MethodPut:
		return "PUT", nil
	case types.MethodTrace:
		return "TRACE", nil
	case types.MethodOther:
		return m.Other(), fmt.Errorf("unknown http method 'other'")
	default:
		return "", fmt.Errorf("failed to convert http method")
	}
}

func toHttpHeader(src types.Fields, dest *http.Header) {
	for _, f := range src.CopyAll() {
		key := f.F0
		value := string(f.F1)
		dest.Add(key, value)
	}
}
