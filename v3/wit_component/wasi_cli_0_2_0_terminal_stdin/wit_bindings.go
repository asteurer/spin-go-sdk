package wasi_cli_0_2_0_terminal_stdin

import (
	"runtime"
	"unsafe"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/wasi_cli_0_2_0_terminal_input"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_types"
)

type TerminalInput = wasi_cli_0_2_0_terminal_input.TerminalInput

//go:wasmimport wasi:cli/terminal-stdin@0.2.0 get-terminal-stdin
func wasm_import_get_terminal_stdin(arg0 uintptr)

func GetTerminalStdin() wit_types.Option[*wasi_cli_0_2_0_terminal_input.TerminalInput] {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, 8, 4))
	wasm_import_get_terminal_stdin(returnArea)
	var option wit_types.Option[*wasi_cli_0_2_0_terminal_input.TerminalInput]
	switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))) {
	case 0:

		option = wit_types.None[*wasi_cli_0_2_0_terminal_input.TerminalInput]()
	case 1:

		option = wit_types.Some[*wasi_cli_0_2_0_terminal_input.TerminalInput](wasi_cli_0_2_0_terminal_input.TerminalInputFromOwnHandle(int32(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(returnArea), 4))))))
	default:
		panic("unreachable")
	}
	result := option
	return result

}
