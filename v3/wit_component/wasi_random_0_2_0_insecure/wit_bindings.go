package wasi_random_0_2_0_insecure

import (
	"runtime"
	"unsafe"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
)

//go:wasmimport wasi:random/insecure@0.2.0 get-insecure-random-bytes
func wasm_import_get_insecure_random_bytes(arg0 int64, arg1 uintptr)

func GetInsecureRandomBytes(len uint64) []uint8 {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, (2 * 4), 4))
	wasm_import_get_insecure_random_bytes(int64(len), returnArea)
	value := unsafe.Slice((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4)))
	result := value
	return result

}

//go:wasmimport wasi:random/insecure@0.2.0 get-insecure-random-u64
func wasm_import_get_insecure_random_u64() int64

func GetInsecureRandomU64() uint64 {

	result := wasm_import_get_insecure_random_u64()
	return uint64(result)

}
