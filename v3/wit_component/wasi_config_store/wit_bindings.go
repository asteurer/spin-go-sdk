package wasi_config_store

import (
	"runtime"
	"unsafe"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_types"
)

const (
	// This indicates an error from an "upstream" config source.
	// As this could be almost _anything_ (such as Vault, Kubernetes ConfigMaps, KeyValue buckets, etc),
	// the error message is a string.
	ErrorUpstream uint8 = 0
	// This indicates an error from an I/O operation.
	// As this could be almost _anything_ (such as a file read, network connection, etc),
	// the error message is a string.
	// Depending on how this ends up being consumed,
	// we may consider moving this to use the `wasi:io/error` type instead.
	// For simplicity right now in supporting multiple implementations, it is being left as a string.
	ErrorIo uint8 = 1
)

// An error type that encapsulates the different errors that can occur fetching configuration values.
type Error struct {
	tag   uint8
	value any
}

func (self Error) Tag() uint8 {
	return self.tag
}

func (self Error) Upstream() string {
	if self.tag != ErrorUpstream {
		panic("tag mismatch")
	}
	return self.value.(string)
}
func (self Error) Io() string {
	if self.tag != ErrorIo {
		panic("tag mismatch")
	}
	return self.value.(string)
}

func MakeErrorUpstream(value string) Error {
	return Error{ErrorUpstream, value}
}
func MakeErrorIo(value string) Error {
	return Error{ErrorIo, value}
}

//go:wasmimport wasi:config/store@0.2.0-draft-2024-09-27 get
func wasm_import_get(arg0 uintptr, arg1 uint32, arg2 uintptr)

func Get(key string) wit_types.Result[wit_types.Option[string], Error] {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, (4 * 4), 4))
	utf8 := unsafe.Pointer(unsafe.StringData(key))
	pinner.Pin(utf8)
	wasm_import_get(uintptr(utf8), uint32(len(key)), returnArea)
	var result wit_types.Result[wit_types.Option[string], Error]
	switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))) {
	case 0:
		var option wit_types.Option[string]
		switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4))) {
		case 0:

			option = wit_types.None[string]()
		case 1:
			value := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			option = wit_types.Some[string](value)
		default:
			panic("unreachable")
		}

		result = wit_types.Ok[wit_types.Option[string], Error](option)
	case 1:
		var variant Error
		switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4))) {
		case 0:
			value0 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			variant = MakeErrorUpstream(value0)

		case 1:
			value1 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			variant = MakeErrorIo(value1)

		default:
			panic("unreachable")
		}

		result = wit_types.Err[wit_types.Option[string], Error](variant)
	default:
		panic("unreachable")
	}
	result2 := result
	return result2

}

//go:wasmimport wasi:config/store@0.2.0-draft-2024-09-27 get-all
func wasm_import_get_all(arg0 uintptr)

func GetAll() wit_types.Result[[]wit_types.Tuple2[string, string], Error] {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, (4 * 4), 4))
	wasm_import_get_all(returnArea)
	var result3 wit_types.Result[[]wit_types.Tuple2[string, string], Error]
	switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))) {
	case 0:
		result := make([]wit_types.Tuple2[string, string], 0, *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4))))
		for index := 0; index < int(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))); index++ {
			base := unsafe.Add(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4)))), index*(4*4))
			value := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(base), 0))))), *(*uint32)(unsafe.Add(unsafe.Pointer(base), 4)))
			value0 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(base), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(base), (3 * 4))))

			result = append(result, wit_types.Tuple2[string, string]{value, value0})
		}

		result3 = wit_types.Ok[[]wit_types.Tuple2[string, string], Error](result)
	case 1:
		var variant Error
		switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4))) {
		case 0:
			value1 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			variant = MakeErrorUpstream(value1)

		case 1:
			value2 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			variant = MakeErrorIo(value2)

		default:
			panic("unreachable")
		}

		result3 = wit_types.Err[[]wit_types.Tuple2[string, string], Error](variant)
	default:
		panic("unreachable")
	}
	result4 := result3
	return result4

}
