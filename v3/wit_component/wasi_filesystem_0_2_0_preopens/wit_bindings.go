package wasi_filesystem_0_2_0_preopens

import (
	"runtime"
	"unsafe"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/wasi_filesystem_0_2_0_types"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_types"
)

type Descriptor = wasi_filesystem_0_2_0_types.Descriptor

//go:wasmimport wasi:filesystem/preopens@0.2.0 get-directories
func wasm_import_get_directories(arg0 uintptr)

func GetDirectories() []wit_types.Tuple2[*wasi_filesystem_0_2_0_types.Descriptor, string] {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, (2 * 4), 4))
	wasm_import_get_directories(returnArea)
	result := make([]wit_types.Tuple2[*wasi_filesystem_0_2_0_types.Descriptor, string], 0, *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4)))
	for index := 0; index < int(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4))); index++ {
		base := unsafe.Add(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0)))), index*(3*4))
		value := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(base), 4))))), *(*uint32)(unsafe.Add(unsafe.Pointer(base), (2 * 4))))

		result = append(result, wit_types.Tuple2[*wasi_filesystem_0_2_0_types.Descriptor, string]{wasi_filesystem_0_2_0_types.DescriptorFromOwnHandle(int32(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(base), 0))))), value})
	}

	result0 := result
	return result0

}
