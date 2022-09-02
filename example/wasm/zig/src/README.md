# Zig wasm serverless function

## Build

```sh
zig build-exe src/main.zig -target wasm32-wasi
cp main.wasm ../sfn.wasm
```
