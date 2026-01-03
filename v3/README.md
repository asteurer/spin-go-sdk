# Overview
The `wasip2` implementation of the Spinframework Go SDK.

## Generating the WIT bindings
Whenever WIT files are changed/added to the `v3/wit` directory, the bindings  in `v3/wit_component` need to be regenerated.

### Prerequisites
- [**componentize-go**](https://github.com/asteurer/componentize-go) - Latest version

### Run
```sh
cd v3
componentize-go -w http-trigger -d ./wit bindings -o wit_component
```
