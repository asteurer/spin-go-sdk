package wit_component

import (
	"runtime"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/export_wasi_http_0_2_0_incoming_handler"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wasi_http_0_2_0_types"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
)

var staticPinner = runtime.Pinner{}
var exportReturnArea = uintptr(wit_runtime.Allocate(&staticPinner, 0, 1))
var syncExportPinner = runtime.Pinner{}

//go:wasmexport wasi:http/incoming-handler@0.2.0#handle
func wasm_export_wasi_http_0_2_0_incoming_handler_handle(arg0 int32, arg1 int32) {
	export_wasi_http_0_2_0_incoming_handler.Exports.Handle(wasi_http_0_2_0_types.IncomingRequestFromOwnHandle(int32(uintptr(arg0))), wasi_http_0_2_0_types.ResponseOutparamFromOwnHandle(int32(uintptr(arg1))))
}
