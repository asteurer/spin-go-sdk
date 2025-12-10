package fermyon_spin_2_0_0_variables

import (
	"runtime"
	"unsafe"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_types"
)

const (
	// The provided variable name is invalid.
	ErrorInvalidName uint8 = 0
	// The provided variable is undefined.
	ErrorUndefined uint8 = 1
	// A variables provider specific error has occurred.
	ErrorProvider uint8 = 2
	// Some implementation-specific error has occurred.
	ErrorOther uint8 = 3
)

// The set of errors which may be raised by functions in this interface.
type Error struct {
	tag   uint8
	value any
}

func (self Error) Tag() uint8 {
	return self.tag
}

func (self Error) InvalidName() string {
	if self.tag != ErrorInvalidName {
		panic("tag mismatch")
	}
	return self.value.(string)
}
func (self Error) Undefined() string {
	if self.tag != ErrorUndefined {
		panic("tag mismatch")
	}
	return self.value.(string)
}
func (self Error) Provider() string {
	if self.tag != ErrorProvider {
		panic("tag mismatch")
	}
	return self.value.(string)
}
func (self Error) Other() string {
	if self.tag != ErrorOther {
		panic("tag mismatch")
	}
	return self.value.(string)
}

func MakeErrorInvalidName(value string) Error {
	return Error{ErrorInvalidName, value}
}
func MakeErrorUndefined(value string) Error {
	return Error{ErrorUndefined, value}
}
func MakeErrorProvider(value string) Error {
	return Error{ErrorProvider, value}
}
func MakeErrorOther(value string) Error {
	return Error{ErrorOther, value}
}

//go:wasmimport fermyon:spin/variables@2.0.0 get
func wasm_import_get(arg0 uintptr, arg1 uint32, arg2 uintptr)

func Get(name string) wit_types.Result[string, Error] {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, (4 * 4), 4))
	utf8 := unsafe.Pointer(unsafe.StringData(name))
	pinner.Pin(utf8)
	wasm_import_get(uintptr(utf8), uint32(len(name)), returnArea)
	var result wit_types.Result[string, Error]
	switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))) {
	case 0:
		value := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4))))

		result = wit_types.Ok[string, Error](value)
	case 1:
		var variant Error
		switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4))) {
		case 0:
			value0 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			variant = MakeErrorInvalidName(value0)

		case 1:
			value1 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			variant = MakeErrorUndefined(value1)

		case 2:
			value2 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			variant = MakeErrorProvider(value2)

		case 3:
			value3 := unsafe.String((*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (2 * 4)))))), *(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), (3 * 4))))

			variant = MakeErrorOther(value3)

		default:
			panic("unreachable")
		}

		result = wit_types.Err[string, Error](variant)
	default:
		panic("unreachable")
	}
	result4 := result
	return result4

}
