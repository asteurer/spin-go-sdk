package wasi_sockets_0_2_0_ip_name_lookup

import (
	"runtime"
	"unsafe"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/wasi_io_0_2_0_poll"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wasi_sockets_0_2_0_network"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_types"
)

type Pollable = wasi_io_0_2_0_poll.Pollable
type Network = wasi_sockets_0_2_0_network.Network
type ErrorCode = wasi_sockets_0_2_0_network.ErrorCode
type IpAddress = wasi_sockets_0_2_0_network.IpAddress

//go:wasmimport wasi:sockets/ip-name-lookup@0.2.0 [resource-drop]resolve-address-stream
func resourceDropResolveAddressStream(handle int32)

type ResolveAddressStream struct {
	handle *wit_runtime.Handle
}

func (self *ResolveAddressStream) TakeHandle() int32 {
	return self.handle.Take()
}

func (self *ResolveAddressStream) Handle() int32 {
	return self.handle.Use()
}

func (self *ResolveAddressStream) Drop() {
	handle := self.handle.TakeOrNil()
	if handle != 0 {
		resourceDropResolveAddressStream(handle)
	}
}

func ResolveAddressStreamFromOwnHandle(handleValue int32) *ResolveAddressStream {
	handle := wit_runtime.MakeHandle(handleValue)
	value := &ResolveAddressStream{handle}
	runtime.AddCleanup(value, func(_ int) {
		handleValue := handle.TakeOrNil()
		if handleValue != 0 {
			resourceDropResolveAddressStream(handleValue)
		}
	}, 0)
	return value
}

func ResolveAddressStreamFromBorrowHandle(handleValue int32) *ResolveAddressStream {
	return ResolveAddressStreamFromOwnHandle(handleValue)
}

//go:wasmimport wasi:sockets/ip-name-lookup@0.2.0 resolve-addresses
func wasm_import_resolve_addresses(arg0 int32, arg1 uintptr, arg2 uint32, arg3 uintptr)

func ResolveAddresses(network *wasi_sockets_0_2_0_network.Network, name string) wit_types.Result[*ResolveAddressStream, wasi_sockets_0_2_0_network.ErrorCode] {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, 8, 4))
	utf8 := unsafe.Pointer(unsafe.StringData(name))
	pinner.Pin(utf8)
	wasm_import_resolve_addresses((network).Handle(), uintptr(utf8), uint32(len(name)), returnArea)
	var result wit_types.Result[*ResolveAddressStream, wasi_sockets_0_2_0_network.ErrorCode]
	switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))) {
	case 0:

		result = wit_types.Ok[*ResolveAddressStream, wasi_sockets_0_2_0_network.ErrorCode](ResolveAddressStreamFromOwnHandle(int32(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(returnArea), 4))))))
	case 1:

		result = wit_types.Err[*ResolveAddressStream, wasi_sockets_0_2_0_network.ErrorCode](uint8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4)))))
	default:
		panic("unreachable")
	}
	result0 := result
	return result0

}

//go:wasmimport wasi:sockets/ip-name-lookup@0.2.0 [method]resolve-address-stream.resolve-next-address
func wasm_import_method_resolve_address_stream_resolve_next_address(arg0 int32, arg1 uintptr)

func (self *ResolveAddressStream) ResolveNextAddress() wit_types.Result[wit_types.Option[wasi_sockets_0_2_0_network.IpAddress], wasi_sockets_0_2_0_network.ErrorCode] {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, 22, 2))
	wasm_import_method_resolve_address_stream_resolve_next_address((self).Handle(), returnArea)
	var result wit_types.Result[wit_types.Option[wasi_sockets_0_2_0_network.IpAddress], wasi_sockets_0_2_0_network.ErrorCode]
	switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))) {
	case 0:
		var option wit_types.Option[wasi_sockets_0_2_0_network.IpAddress]
		switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 2))) {
		case 0:

			option = wit_types.None[wasi_sockets_0_2_0_network.IpAddress]()
		case 1:
			var variant wasi_sockets_0_2_0_network.IpAddress
			switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4))) {
			case 0:

				variant = wasi_sockets_0_2_0_network.MakeIpAddressIpv4(wit_types.Tuple4[uint8, uint8, uint8, uint8]{uint8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 6)))), uint8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 7)))), uint8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 8)))), uint8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 9))))})

			case 1:

				variant = wasi_sockets_0_2_0_network.MakeIpAddressIpv6(wit_types.Tuple8[uint16, uint16, uint16, uint16, uint16, uint16, uint16, uint16]{uint16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 6)))), uint16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 8)))), uint16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 10)))), uint16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 12)))), uint16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 14)))), uint16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 16)))), uint16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 18)))), uint16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 20))))})

			default:
				panic("unreachable")
			}

			option = wit_types.Some[wasi_sockets_0_2_0_network.IpAddress](variant)
		default:
			panic("unreachable")
		}

		result = wit_types.Ok[wit_types.Option[wasi_sockets_0_2_0_network.IpAddress], wasi_sockets_0_2_0_network.ErrorCode](option)
	case 1:

		result = wit_types.Err[wit_types.Option[wasi_sockets_0_2_0_network.IpAddress], wasi_sockets_0_2_0_network.ErrorCode](uint8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 2)))))
	default:
		panic("unreachable")
	}
	result0 := result
	return result0

}

//go:wasmimport wasi:sockets/ip-name-lookup@0.2.0 [method]resolve-address-stream.subscribe
func wasm_import_method_resolve_address_stream_subscribe(arg0 int32) int32

func (self *ResolveAddressStream) Subscribe() *wasi_io_0_2_0_poll.Pollable {

	result := wasm_import_method_resolve_address_stream_subscribe((self).Handle())
	return wasi_io_0_2_0_poll.PollableFromOwnHandle(int32(uintptr(result)))

}
