module github.com/spinframework/spin-go-sdk/v4

go 1.25.5

require (
	github.com/bytecodealliance/wit-bindgen v0.0.0-00010101000000-000000000000
	github.com/julienschmidt/httprouter v1.3.0
	github.com/spinframework/spin-go-sdk/v4/wit_component v0.1.0
)

replace github.com/spinframework/spin-go-sdk/v4/wit_component v0.1.0 => ./wit_component

replace github.com/bytecodealliance/wit-bindgen => ../../wit-bindgen/crates/go/src/package
