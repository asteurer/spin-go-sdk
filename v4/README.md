# Overview
The `wasip3` implementation of the Spinframework Go SDK.

> [!Caution]
> This is a work in progress

## Generating the WIT bindings
Whenever WIT files are changed/added to the `v4/wit` directory, the bindings  in `v4/wit_component` need to be regenerated.

### Prerequisites
- [**componentize-go**](https://github.com/asteurer/componentize-go) - Latest version

### Run
```sh
cd v3
componentize-go -w http-trigger -d ./wit bindings -o wit_component
```
