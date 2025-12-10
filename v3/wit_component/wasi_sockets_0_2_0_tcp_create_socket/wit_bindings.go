package wasi_sockets_0_2_0_tcp_create_socket

import (
	"runtime"
	"unsafe"

	"github.com/spinframework/spin-go-sdk/v3/wit_component/wasi_sockets_0_2_0_network"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wasi_sockets_0_2_0_tcp"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_runtime"
	"github.com/spinframework/spin-go-sdk/v3/wit_component/wit_types"
)

type Network = wasi_sockets_0_2_0_network.Network
type ErrorCode = wasi_sockets_0_2_0_network.ErrorCode
type IpAddressFamily = wasi_sockets_0_2_0_network.IpAddressFamily
type TcpSocket = wasi_sockets_0_2_0_tcp.TcpSocket

//go:wasmimport wasi:sockets/tcp-create-socket@0.2.0 create-tcp-socket
func wasm_import_create_tcp_socket(arg0 int32, arg1 uintptr)

func CreateTcpSocket(addressFamily wasi_sockets_0_2_0_network.IpAddressFamily) wit_types.Result[*wasi_sockets_0_2_0_tcp.TcpSocket, wasi_sockets_0_2_0_network.ErrorCode] {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	returnArea := uintptr(wit_runtime.Allocate(pinner, 8, 4))
	wasm_import_create_tcp_socket(int32(addressFamily), returnArea)
	var result wit_types.Result[*wasi_sockets_0_2_0_tcp.TcpSocket, wasi_sockets_0_2_0_network.ErrorCode]
	switch uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 0))) {
	case 0:

		result = wit_types.Ok[*wasi_sockets_0_2_0_tcp.TcpSocket, wasi_sockets_0_2_0_network.ErrorCode](wasi_sockets_0_2_0_tcp.TcpSocketFromOwnHandle(int32(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(returnArea), 4))))))
	case 1:

		result = wit_types.Err[*wasi_sockets_0_2_0_tcp.TcpSocket, wasi_sockets_0_2_0_network.ErrorCode](uint8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(returnArea), 4)))))
	default:
		panic("unreachable")
	}
	result0 := result
	return result0

}
