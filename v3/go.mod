module github.com/spinframework/spin-go-sdk/v3

go 1.25

require (
	github.com/julienschmidt/httprouter v1.3.0
	github.com/spinframework/spin-go-sdk/v3/wit_component v0.1.0
)

replace github.com/spinframework/spin-go-sdk/v3/wit_component v0.1.0 => ./wit_component
