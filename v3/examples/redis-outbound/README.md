# Requirements
# Requirements
- [**go**](https://go.dev/dl/) - v1.25+
- [**spin**](https://github.com/spinframework/spin) - Latest version
- [**wasm-tools**](https://github.com/bytecodealliance/wasm-tools) - Latest version
- [**just**](https://github.com/casey/just) - Command runner

# Usage

In one terminal window, you'll run a Redis container:
```sh
docker run -p 6379:6379 redis:8.2
```

In another terminal, you'll run your Spin app:
```sh
just build
spin up
```

In yet another terminal, you'll interact with the Spin app:
```sh
curl localhost:3000
```
