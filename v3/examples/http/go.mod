module github.com/spinframework/spin-go-sdk/v3/examples/http

go 1.25

require (
	github.com/spinframework/spin-go-sdk/v3 v3.0.0
	github.com/spinframework/spin-go-sdk/v3/wit_component v0.1.0
)

require github.com/julienschmidt/httprouter v1.3.0 // indirect

replace github.com/spinframework/spin-go-sdk/v3 => ../../
replace github.com/spinframework/spin-go-sdk/v3/wit_component => ../../wit_component
