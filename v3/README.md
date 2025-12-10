# Overview
The `wasip2` implementation of the Spinframework Go SDK.

## Generating the WIT bindings
### Prerequisites
- [**wit-bindgen**](https://github.com/bytecodealliance/wit-bindgen) - Latest version
- [**go**](https://go.dev/dl/) - v1.25+
- [**spin**](https://github.com/spinframework/spin) - Latest version

### Run
```sh
cd v3
wit-bindgen go -w http-trigger --out-dir wit_component ./wit
```

## Running the HTTP example application
### Prerequisites
- [**go**](https://go.dev/dl/) - v1.25+
- [**spin**](https://github.com/spinframework/spin) - Latest version
- [**wasm-tools**](https://github.com/bytecodealliance/wasm-tools) - v1.239.0
- [**just**](https://github.com/casey/just) - Command runner

## Run
```sh
cd v3/examples/http
just build
spin up

# In a different terminal...
curl localhost:3000/hello
```
