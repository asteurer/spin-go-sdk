package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bytecodealliance/wit-bindgen/wit_types"
	types "github.com/spinframework/spin-go-sdk/v4/wit_component/wasi_http_0_3_0_rc_2025_09_16_types"
)

var _ http.ResponseWriter = &responseWriter{}

type responseWriter struct {
	response    *types.Response
	httpHeaders http.Header
	wasiHeaders types.Fields
	statusCode  int
	reconciled  bool
	bodyWriter  *wit_types.StreamWriter[uint8]
}

func NewHttpResponseWriter(tx *wit_types.StreamWriter[uint8]) *responseWriter {
	return &responseWriter{
		httpHeaders: http.Header{},
		wasiHeaders: *types.MakeFields(),
		bodyWriter:  tx,
	}
}

func (rw *responseWriter) Header() http.Header {
	return rw.httpHeaders
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if err := rw.reconcile(); err != nil {
		return 0, err
	}

	rw.bodyWriter.WriteAll(data)
	return len(data), nil
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}

func (rw *responseWriter) reconcile() error {
	if rw.reconciled {
		return nil
	}

	if err := rw.reconcileHeaders(); err != nil {
		return err
	}

	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}

	// TODO: Unsure if this needs to be replicated:
	// // setting any headers after this will cause panic
	// row.response = *types.MakeOutgoingResponse(&row.wasiHeaders)

	rw.response.SetStatusCode(types.StatusCode(rw.statusCode))

	// Mark that reconcilliation has already happened
	rw.reconciled = true

	return nil
}

func (rw *responseWriter) reconcileHeaders() error {
	// Get all WASI headers to track any deletions
	// TODO: This needs to be tested
	existingHeaders := make(map[string]bool)
	for _, entry := range rw.wasiHeaders.CopyAll() {
		existingHeaders[entry.F0] = true
	}

	for key, vals := range rw.httpHeaders {
		fieldVals := []types.FieldValue{}
		for _, val := range vals {
			fieldVals = append(fieldVals, types.FieldValue(val))
		}

		if result := rw.wasiHeaders.Set(key, fieldVals); result.IsErr() {
			switch result.Err().Tag() {
			case types.HeaderErrorInvalidSyntax:
				return fmt.Errorf("failed to set header %s to [%s]: invalid syntax", key, strings.Join(vals, ","))
			case types.HeaderErrorForbidden:
				return fmt.Errorf("failed to set forbidden header key %s", key)
			case types.HeaderErrorImmutable:
				return fmt.Errorf("failed to set header on immutable header fields")
			default:
				return fmt.Errorf("not sure what happened here?")
			}
		}

		// Mark as processed
		delete(existingHeaders, key)
	}

	// Delete headers that were removed from httpHeaders
	for headerName := range existingHeaders {
		if result := rw.wasiHeaders.Delete(headerName); result.IsErr() {
			switch result.Err().Tag() {
			case types.HeaderErrorInvalidSyntax:
				return fmt.Errorf("failed to delete header %s: invalid syntax", headerName)
			case types.HeaderErrorForbidden:
				return fmt.Errorf("failed to delete forbidden header key %s", headerName)
			case types.HeaderErrorImmutable:
				return fmt.Errorf("failed to delete header on immutable header fields")
			default:
				return fmt.Errorf("not sure what happened here?")
			}
		}
	}

	return nil
}

func (rw *responseWriter) getHeaders() *types.Fields {
	result := types.FieldsFromList(rw.wasiHeaders.CopyAll())
	if result.IsOk() {
		return result.Ok()
	} else {
		return nil
	}
}
