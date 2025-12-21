# Overview
The `wasip2` implementation of the Spinframework Go SDK.

## Generating the WIT bindings
Whenever WIT files are changed/added to the `v3/wit` directory, the bindings  in `v3/wit_component` need to be regenerated.

### Prerequisites
- [**wit-bindgen**](https://github.com/bytecodealliance/wit-bindgen) - Latest version

### Run
```sh
cd v3
wit-bindgen go -w http-trigger --out-dir wit_component ./wit
```
