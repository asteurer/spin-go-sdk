package wasi_random_0_2_0_random

import (
	"runtime"
	"unsafe"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
)

//go:wasmimport wasi:random/random@0.2.0 get-random-bytes
func wasm_import_get_random_bytes(arg0 int64, arg1 uintptr)

func GetRandomBytes(len uint64) []uint8 {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, (2 * 4), 4))
	wasm_import_get_random_bytes(int64(len), returnArea)
	value := unsafe.Slice((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4)))
	result := value
	return result

}

//go:wasmimport wasi:random/random@0.2.0 get-random-u64
func wasm_import_get_random_u64() int64

func GetRandomU64() uint64 {

	result := wasm_import_get_random_u64()
	return uint64(result)

}
