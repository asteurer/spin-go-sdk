module github.com/spinframework/spin-go-sdk/v3/http/testdata/http

go 1.25.5

require github.com/spinframework/spin-go-sdk/v3 v3.0.0

require (
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	go.bytecodealliance.org/pkg v0.2.1 // indirect
)

replace github.com/spinframework/spin-go-sdk/v3 => ../../../

replace github.com/bytecodealliance/wit-bindgen => github.com/bytecodealliance/wit-bindgen/crates/go/src/package v0.51.0
